/* Copyright (C) Fedir Petryk */

package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
	appMIddleware "github.com/serajam/realestate-sample-app/internal/adapters/handlers/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}

	return nil
}

func NewServer(logger *zap.Logger) *echo.Echo {
	srv := echo.New()
	srv.Validator = &CustomValidator{validator: validator.New()}
	srv.Use(
		appMIddleware.ZapLoggerWithConfig(
			logger, appMIddleware.ZapLoggerConfig{
				Skipper: func(c echo.Context) bool {
					return !strings.HasPrefix(c.Request().RequestURI, "/api")
				},
			},
		),
	)
	srv.Use(middleware.CORS())
	srv.Use(
		middleware.GzipWithConfig(
			middleware.GzipConfig{
				Level: 5,
			},
		),
	)

	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	srv.Use(middleware.RateLimiterWithConfig(config))
	srv.Use(
		middleware.RecoverWithConfig(
			middleware.RecoverConfig{
				StackSize: 1 << 10, // 1 KB
				LogLevel:  log.ERROR,
			},
		),
	)

	srv.HTTPErrorHandler = common.ErrorHandler(logger.Sugar())

	return srv
}
