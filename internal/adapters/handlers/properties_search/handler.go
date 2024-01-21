/* Copyright (C) Fedir Petryk */

// Package properties_search
// @title  properties actions
// @host   localhost:8090
package properties_search

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (t Handler) AddRoutes(e *echo.Group, m ...echo.MiddlewareFunc) {
	e.POST("/v1/properties/search", t.Search)
	e.POST("/v1/properties/markers", t.SearchMarkers)
	e.GET("/v1/properties/:id/similar", t.Similar)
}
