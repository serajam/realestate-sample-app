/* Copyright (C) Fedir Petryk */

package app

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type postgresConfig struct {
	PostgresPwd  string `env:"POSTGRES_PASSWORD,notEmpty"`
	PostgresUser string `env:"POSTGRES_USER,notEmpty"`
	PostgresHost string `env:"POSTGRES_HOST,notEmpty"`
	PostgresPort string `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresDb   string `env:"POSTGRES_DB,notEmpty"`
	PostgresSsl  string `env:"POSTGRES_SSL" envDefault:"disable"`
}

type redisConfig struct {
	RedisPwd  string `env:"REDIS_PASSWORD,notEmpty"`
	RedisHost string `env:"REDIS_HOST,notEmpty" envDefault:"localhost"`
	RedisPort string `env:"REDIS_PORT" envDefault:"6379"`
}

type natsConfig struct {
	NatsHost string `env:"NATS_HOST,notEmpty"`
	NatsPort string `env:"NATS_PORT" envDefault:"4222"`
}

type awsConfig struct {
	AwsAccessKey string `env:"AWS_KEY,notEmpty"`
	AwsSecretKey string `env:"AWS_SECRET,notEmpty"`
	AwsEndpoint  string `env:"AWS_ENDPOINT,notEmpty"`
	AwsEnableSSL bool   `env:"AWS_SSL" envDefault:"false"`
	AwsBucket    string `env:"AWS_BUCKET" envDefault:"property-images"`
	AwsLocation  string `env:"AWS_LOCATION" envDefault:"us-east-1"`
}

func (nc natsConfig) Url() string {
	return fmt.Sprintf("nats://%s:%s", nc.NatsHost, nc.NatsPort)
}

func (c redisConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

type userConfig struct {
	RegistrationTokenExpiration int    `env:"REGISTRATION_TOKEN_EXPIRATION" envDefault:"24"` // hours
	PwdResetExpiration          int    `env:"PWD_RESET_EXPIRATION" envDefault:"1"`           // hours
	PwdResetUrl                 string `env:"PWD_RESET_URL" envDefault:"http://localhost:8080/api/v1/set-password"`
}

type authConfig struct {
	AccessTokenTTL     int64  `env:"ACCESS_TOKEN_TTL" envDefault:"15"`     // minutes
	RefreshTokenTTL    int64  `env:"REFRESH_TOKEN_TTL" envDefault:"10080"` // minutes
	AccessTokenSecret  string `env:"ACCESS_TOKEN_SECRET" envDefault:"dummy"`
	RefreshTokenSecret string `env:"REFRESH_TOKEN_SECRET" envDefault:"dummy"`
}

type emailConfig struct {
	SmtpHost   string `env:"SMTP_HOST" envDefault:"localhost"`
	SmtpPort   string `env:"SMTP_PORT" envDefault:"1025"`
	SmtpLogin  string `env:"SMTP_LOGIN"`
	SmtpPwd    string `env:"SMTP_PWD"`
	SmtpSender string `env:"SMTP_SENDER"`
}

type propConfig struct {
	PropertiesSearchLimit int `env:"PROPERTIES_SEARCH_LIMIT" envDefault:"50"`
}

type telebotConfig struct {
	Token     string `env:"TELEGRAM_TOKEN"`
	ChannelId int64  `env:"TELEGRAM_CHANNEL_ID"`
}

// @TODO migrate to toml or yaml
type imageProcessingConfig struct {
}

func (c postgresConfig) Url() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.PostgresUser,
		c.PostgresPwd,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDb,
		c.PostgresSsl,
	)
}

type appConfig struct {
	userConfig
	emailConfig
	propConfig
	redisConfig
	authConfig
	natsConfig
	awsConfig
	telebotConfig

	postgresUrl string

	Port int `env:"APP_PORT" envDefault:"8080"`

	PostgresEnableQueryDebug bool `env:"POSTGRES_QUERY_DEBUG" envDefault:"true"`
	GenerateTestData         bool `env:"GENERATE_TEST_DATA" envDefault:"false"`
	CleanupDB                bool `env:"CLEANUP_DB" envDefault:"false"`
}

func initConfig() (*appConfig, error) {
	appCfg := appConfig{}

	cfgDb := postgresConfig{}
	if err := env.Parse(&cfgDb); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into postgresConfig")
	}

	cfgRedis := redisConfig{}
	if err := env.Parse(&cfgRedis); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into redisConfig")
	}

	cfgNats := natsConfig{}
	if err := env.Parse(&cfgNats); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into natsConfig")
	}

	cfgAWS := awsConfig{}
	if err := env.Parse(&cfgAWS); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into awsConfig")
	}

	userCfg := userConfig{RegistrationTokenExpiration: 24}
	if err := env.Parse(&userCfg); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into userCfg")
	}

	authCfg := authConfig{}
	if err := env.Parse(&authCfg); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into userCfg")
	}

	emailCfg := emailConfig{}
	if err := env.Parse(&emailCfg); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into emailCfg")
	}

	propertiesCfg := propConfig{}
	if err := env.Parse(&emailCfg); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into propertiesCfg")
	}

	teleCfg := telebotConfig{}
	if err := env.Parse(&teleCfg); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into teleCfg")
	}

	if err := env.Parse(&appCfg); err != nil {
		return nil, errors.Wrap(err, "initConfig: error parsing env data into appConfig")
	}

	appCfg.postgresUrl = cfgDb.Url()
	appCfg.userConfig = userCfg
	appCfg.emailConfig = emailCfg
	appCfg.propConfig = propertiesCfg
	appCfg.redisConfig = cfgRedis
	appCfg.natsConfig = cfgNats
	appCfg.authConfig = authCfg
	appCfg.awsConfig = cfgAWS
	appCfg.telebotConfig = teleCfg

	return &appCfg, nil
}
