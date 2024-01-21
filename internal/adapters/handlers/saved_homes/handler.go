/* Copyright (C) Fedir Petryk */

// Package saved_homes
// @title  properties actions
// @host   localhost:8090
package saved_homes

import (
	"github.com/labstack/echo/v4"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

type Handler struct {
	common.BaseHandler

	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (t Handler) AddRoutes(e *echo.Group, m ...echo.MiddlewareFunc) {
	e.PUT("/v1/saved-homes/:id", t.Create, m...)
	e.DELETE("/v1/saved-homes/:id", t.Delete, m...)
	e.POST("/v1/saved-homes/list", t.List, m...)
}
