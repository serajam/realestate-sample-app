/* Copyright (C) Fedir Petryk */

// Package search_filters
// @title  searchfilters actions
// @host   localhost:8090
package search_filters

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
	e.GET("/v1/saved-searches/:id", t.Get, m...)
	e.DELETE("/v1/saved-searches/:id", t.Delete, m...)
	e.POST("/v1/saved-searches", t.Create, m...)
	e.PUT("/v1/saved-searches/:id", t.Update, m...)
	e.GET("/v1/saved-searches", t.List, m...)
}
