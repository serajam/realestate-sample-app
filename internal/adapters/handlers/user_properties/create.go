/* Copyright (C) Fedir Petryk */

package user_properties

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Create
//
// @Summary Create propertyCreateRequest
// @Description Create propertyCreateRequest
// @Tags properties
// @Accept json
// @Produce json
// @Param propertyCreate body propertyCreateRequest true "property data"
// @Success 200 {object} common.DefaultResponse{data=propertyResponse}
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/user/properties  [POST]
func (t Handler) Create(ctx echo.Context) error {
	req := &propertyCreateRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrorResponse{Message: err.Error()})
	}

	ok, err := t.ValidateRequest(ctx, req)
	if err != nil || !ok {
		return err
	}

	property, err := createPropertyToDomain(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
	}

	property, err = t.userPropService.CreateUserProperty(ctx.Request().Context(), property)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if property == nil {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(http.StatusOK, common.DefaultResponse{Data: propertyToResponse(property)})
}
