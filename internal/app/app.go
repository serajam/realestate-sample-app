/* Copyright (C) Fedir Petryk  */

package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/user_properties"
	properties2 "github.com/serajam/realestate-sample-app/internal/core/domain/properties"

	"github.com/nats-io/nats.go"
	"github.com/serajam/realestate-sample-app/internal/adapters/datastore/storage"
	"github.com/serajam/realestate-sample-app/internal/adapters/external/email"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/properties_search"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/saved_homes"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/search_filters"
	"github.com/serajam/realestate-sample-app/internal/adapters/publishers/emails/user"
	userSub "github.com/serajam/realestate-sample-app/internal/adapters/subscribers/emails"
	"github.com/serajam/realestate-sample-app/internal/core/domain/auth"
	"github.com/serajam/realestate-sample-app/internal/core/services/emails"
	propServices "github.com/serajam/realestate-sample-app/internal/core/services/properties"
	userServices "github.com/serajam/realestate-sample-app/internal/core/services/users"
	"github.com/serajam/realestate-sample-app/internal/infrastructure/notifications"
	"github.com/serajam/realestate-sample-app/internal/infrastructure/pubsub"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"

	_ "github.com/serajam/realestate-sample-app/docs/swagger" // register swagger docs for wrapHandler
	"github.com/serajam/realestate-sample-app/internal/adapters/datastore/repositories"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/middleware"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/properties"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/users"
	"github.com/serajam/realestate-sample-app/internal/infrastructure/datastore"
	"github.com/serajam/realestate-sample-app/internal/test/data"
)

type App struct {
	logger        *zap.SugaredLogger
	server        *echo.Echo
	nats          *nats.Conn
	subscriptions []Subscriber
	config        *appConfig
}

func NewApp(ctx context.Context, logger *zap.Logger) (*App, error) {
	cfg, err := initConfig()
	if err != nil {
		return nil, errors.Wrap(err, "app")
	}

	dbConn, err := datastore.NewPostgresConn(ctx, cfg.postgresUrl, cfg.PostgresEnableQueryDebug, logger.Sugar().With("init", "postgres"))
	if err != nil {
		return nil, errors.Wrap(err, "app")
	}

	redisClient, err := datastore.NewRedisConn(ctx, cfg.Address(), cfg.RedisPwd, logger.Sugar().With("init", "redis"))
	if err != nil {
		return nil, errors.Wrap(err, "app")
	}

	nc, err := pubsub.NewNatsConn(cfg.natsConfig.Url(), logger.Sugar().With("init", "nats"))
	if err != nil {
		return nil, errors.Wrap(err, "app")
	}

	err = datastore.ApplyMigrations(ctx, dbConn, cfg.CleanupDB, logger.Sugar().With("init", "migrations"))
	if err != nil {
		return nil, errors.Wrap(err, "app")
	}

	emailer := email.NewSMTP(cfg.SmtpHost, cfg.SmtpPort, cfg.SmtpLogin, cfg.SmtpPwd, cfg.SmtpSender)

	minioClient, err := datastore.NewMinioConn(
		ctx,
		cfg.AwsEndpoint,
		cfg.AwsAccessKey,
		cfg.AwsSecretKey,
		cfg.AwsEnableSSL,
		logger.Sugar().With("init", "minio"),
	)

	if err != nil {
		return nil, errors.Wrap(err, "app")
	}

	imageStorage, err := storage.NewImagesStorage(ctx, minioClient, cfg.AwsBucket, cfg.AwsLocation)
	if err != nil {
		return nil, errors.Wrap(err, "app")
	}

	dimensions := []properties2.Dimensions{
		{Width: 1920, Height: 1080, ImageType: propServices.ImageTypeOrigin},
		{Width: 290, Height: 170, ImageType: propServices.ImageTypeThumbnail},
	}
	imgSrv := propServices.NewImagesProcessService(dimensions, []string{"jpeg", "png", "jpg"})

	propRepo := repositories.NewProperty(dbConn)
	imagesRepo := repositories.NewPropertyImages(dbConn)
	imgPropService := propServices.NewImageService(imageStorage, imagesRepo, imgSrv, logger.Sugar().With("service", "images"))

	propSearchRepo := repositories.NewPropertySearch(dbConn)
	propService := propServices.NewPropertySrv(
		propRepo,
		propSearchRepo,
		logger.Sugar().With("service", "property"),
	)

	userEmailPublisher := user.NewPublisher(nc, logger.Sugar().With("publisher", "user"))
	usersRepo := repositories.NewUser(dbConn)
	profileService := userServices.NewProfileSrv(
		usersRepo,
		propRepo,
		userEmailPublisher,
		logger.Sugar().With("service", "profile"),
	)

	userPropService := propServices.NewUserPropertySrv(
		propRepo, propSearchRepo, cfg.PropertiesSearchLimit,
		logger.Sugar().With("service", "user_property"),
	)

	propSearchSimilarRepo := repositories.NewSimilarPropertySearch(dbConn)
	propSearchService := propServices.NewPropertySearchSrv(
		propRepo,
		propSearchRepo,
		propSearchSimilarRepo,
		cfg.PropertiesSearchLimit,
		logger.Sugar().With("service", "property_search"),
	)

	regService := userServices.NewSignupSrv(
		usersRepo,
		userEmailPublisher,
		cfg.RegistrationTokenExpiration,
		logger.Sugar().With("service", "signup"),
	)

	accessTokenGen := auth.NewTokenGenerator(cfg.AccessTokenSecret, cfg.AccessTokenTTL)
	refreshTokenGen := auth.NewTokenGenerator(cfg.RefreshTokenSecret, cfg.RefreshTokenTTL)

	tokensRepo := repositories.NewUserTokenAction(dbConn)
	authService := userServices.NewAuthSrv(usersRepo, redisClient, accessTokenGen, refreshTokenGen, logger.Sugar().With("service", "auth"))
	pwdService := userServices.NewPasswordSrv(
		usersRepo, tokensRepo, emailer, cfg.PwdResetExpiration, cfg.PwdResetUrl, logger.Sugar().With("service", "password"),
	)

	searchFiltersRepo := repositories.NewSearchFilters(dbConn)
	searchFiltersService := propServices.NewSearchSrv(searchFiltersRepo, logger.Sugar().With("service", "search_filters"))

	savedHomesRepo := repositories.NewSavedHomes(dbConn)
	savedHomesService := propServices.NewSavedHomesSrv(
		savedHomesRepo,
		propRepo,
		propSearchRepo,
		logger.Sugar().With("service", "saved_homes"),
	)

	if cfg.GenerateTestData {
		go func() {
			data.GenProperties(propRepo, usersRepo, logger.Sugar())
		}()
	}

	userEmailsService := emails.NewUserEmailsSrv(emailer, logger.Sugar().With("service", "user_emails"))
	userSubscriptions, err := userSub.NewSubscriber(nc, userEmailsService, logger.Sugar())
	if err != nil {
		return nil, errors.Wrap(err, "app")
	}

	subscribers := []Subscriber{userSubscriptions}

	server := handlers.NewServer(logger)
	api := server.Group("/api")

	authMdwr := middleware.Auth(authService)

	apiHandlers := []Handler{
		properties.NewHandler(propService, profileService),
		user_properties.NewHandler(imgPropService, userPropService),
		search_filters.NewHandler(searchFiltersService),
		saved_homes.NewHandler(savedHomesService),
		properties_search.NewHandler(propSearchService),
		users.NewHandler(authService, regService, pwdService, profileService),
	}

	for _, h := range apiHandlers {
		h.AddRoutes(api, authMdwr)
	}

	return &App{logger: logger.Sugar(), server: server, config: cfg, nats: nc, subscriptions: subscribers}, nil
}

func (app App) Run() error {
	telebot, err := notifications.NewTelebot(app.config.telebotConfig.Token, app.config.telebotConfig.ChannelId)
	if err != nil {
		app.logger.Errorw("start telebot", "error", err)
	}

	// add swagger and health
	app.server.GET("/swagger/*", echoSwagger.WrapHandler)
	app.server.GET(
		"/health", func(c echo.Context) error {
			return c.JSON(200, map[string]string{"status": "ok"})
		},
	)

	// Start server
	go func() {
		if err := app.server.Start(fmt.Sprintf(":%d", app.config.Port)); err != nil && !errors.Is(
			err, http.ErrServerClosed,
		) {
			app.logger.Errorw("start server", "error", err)
		}
	}()

	app.logger.Infow("Started application", "Port", app.config.Port)

	// Wait for interrupt signal to gracefully shutdown the server with app timeout of 15 seconds.
	// Use app buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err = telebot.Send("app restart")
	if err != nil {
		app.logger.Errorw("send tele notification message", "error", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	for _, sub := range app.subscriptions {
		sub.Shutdown()
	}

	err = app.nats.Drain()
	if err != nil {
		app.logger.Errorw("nats drain", "error", err)
	}

	if err := app.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}

	return nil
}
