/* Copyright (C) Fedir Petryk */

// Package user_properties
// @title  user properties actions
// @host   localhost:8090
package user_properties

import (
	"github.com/labstack/echo/v4"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

type Handler struct {
	userPropService UserPropertiesService
	imgService      ImagesService

	common.BaseHandler
}

func NewHandler(imgService ImagesService, userPropService UserPropertiesService) Handler {
	return Handler{
		imgService:      imgService,
		userPropService: userPropService,
	}
}

func (t Handler) AddRoutes(e *echo.Group, m ...echo.MiddlewareFunc) {
	e.GET("/v1/user/properties", t.ListByUser, m...)
	e.GET("/v1/user/properties/:id", t.GetUserProp, m...)
	e.DELETE("/v1/user/properties/:id", t.Delete, m...)
	e.POST("/v1/user/properties", t.Create, m...)
	e.PUT("/v1/user/properties/:id", t.Update, m...)
	e.POST("/v1/user/properties/:id/image", t.Upload, m...)
	e.PUT("/v1/user/properties/:id/activate", t.Activate, m...)
	e.PUT("/v1/user/properties/:id/deactivate", t.Deactivate, m...)
	e.DELETE("/v1/user/properties/:property_id/image/:image_id", t.DeleteImage, m...)
}
