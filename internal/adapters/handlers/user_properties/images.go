/* Copyright (C) Fedir Petryk */

package user_properties

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Upload
//
// @Summary Create property image
// @Description Create property image
// @Tags properties
// @Accept json
// @Produce json
// @Param imageContent body image true "image"
// @Success 204 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/user/properties/{id}/image  [POST]
func (t Handler) Upload(ctx echo.Context) error {
	// @TODO
	// validation, nginx max body size, etc for this route or content type
	// retrieving image by id and type
	// retrieving images with property

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: err.Error()})
	}

	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return errors.New("invalid user id")
	}

	_, err = t.userPropService.GetUserProperty(ctx.Request().Context(), id, userID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		return err
	}

	err = t.imgService.Save(ctx.Request().Context(), file, userID, id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}

// DeleteImage
//
// @Summary Download property image
// @Description Download property image
// @Tags properties
// @Accept json
// @Success 204 {object} common.EmptyResponse
// @Param property_id path int true "property id"
// @Param image_id path string true "image uuid from property images list"
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/user/properties/{property_id}/image/{image_id}  [DELETE]
func (t Handler) DeleteImage(ctx echo.Context) error {
	id := ctx.Param(common.ImageId)
	imgUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: err.Error()})
	}

	userID, err := common.UserID(ctx)
	if err != nil {
		return err
	}

	propertyID, err := common.IDFromPath(ctx, common.PropertyId)
	if err != nil {
		return err
	}

	_, err = t.userPropService.GetUserProperty(ctx.Request().Context(), propertyID, userID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	err = t.imgService.DeleteOne(ctx.Request().Context(), propertyID, imgUUID.String())
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}
