/* Copyright (C) Fedir Petryk */

package main

import (
	"context"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/serajam/realestate-sample-app/internal/app"
)

var build = "dev"

// @title realestate
// @version 1.0
// @description realestate
// @host      localhost:8080
// @BasePath  /api
// @name realestate
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config := zap.NewDevelopmentConfig()
	// if you're using console encoding, the FunctionKey value can be any
	// non-empty string because console encoding does not print the key.
	config.EncoderConfig.FunctionKey = "F"
	config.DisableStacktrace = true
	config.Encoding = "console"

	logger, _ := config.Build()
	logger = logger.WithOptions(zap.AddCaller())
	logger.Info("Test Logging")

	defer logger.Sync()
	config.EncoderConfig.FunctionKey = "F"

	logger = logger.With(zap.String("build", build))
	logger.Info("Starting application...")

	err := godotenv.Load()
	if err != nil {
		logger.Warn(err.Error())
	}

	app, err := app.NewApp(context.Background(), logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	if err := app.Run(); err != nil {
		logger.Fatal(err.Error())
	}
}
