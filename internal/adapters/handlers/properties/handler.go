/* Copyright (C) Fedir Petryk */

// Package properties
// @title  properties actions
// @host   localhost:8090
package properties

import (
	"github.com/labstack/echo/v4"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

type Handler struct {
	propService    PropertiesService
	profileService ProfileService

	common.BaseHandler
}

func NewHandler(propService PropertiesService, profileService ProfileService) common.Handler {
	return Handler{
		propService:    propService,
		profileService: profileService,
	}
}

func (t Handler) AddRoutes(e *echo.Group, m ...echo.MiddlewareFunc) {
	e.GET("/v1/properties/:id", t.Get)
	e.POST("/v1/properties/list", t.List)
}
