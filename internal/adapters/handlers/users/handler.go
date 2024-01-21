/* Copyright (C) Fedir Petryk */

// Package users
// @title  users actions
// @host   localhost:8090
package users

import (
	"github.com/labstack/echo/v4"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

type Handler struct {
	authService     AuthService
	signinService   SignUpService
	pwdResetService ResetPwdService
	profileService  ProfileService

	common.BaseHandler
}

func NewHandler(
	authService AuthService, signinService SignUpService, pwdResetService ResetPwdService,
	profileService ProfileService,
) Handler {
	return Handler{
		profileService: profileService, authService: authService, signinService: signinService,
		pwdResetService: pwdResetService,
	}
}

func (h Handler) AddRoutes(e *echo.Group, m ...echo.MiddlewareFunc) {
	e.POST("/v1/sign-up", h.SignUp)
	e.POST("/v1/sign-in", h.SignIn)
	e.DELETE("/v1/sign-out", h.SignOut, m...)
	e.DELETE("/v1/sign-out/all", h.SignOutAll, m...)
	e.PUT("/v1/refresh-token", h.Refresh, m...)
	e.POST("/v1/reset-password", h.ResetPwd)
	e.PUT("/v1/set-password/:token", h.NewPwd)

	e.GET("/v1/user/profile", h.GetProfile, m...)
	e.PATCH("/v1/user/profile", h.PutProfile, m...)
	e.PUT("/v1/user/profile/deactivate", h.Deactivate, m...)
}
