/* Copyright (C) Fedir Petryk */

package properties_search

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Search
//
// @Summary Search properties.
// @Description Search properties using different filters
// @Tags properties
// @Accept json
// @Produce json
// @Param Search body propertySearchRequest true "property search body"
// @Success 200 {object} common.DefaultResponse{data=[]propertyResponse}
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/properties/search [POST]
func (t Handler) Search(ctx echo.Context) error {
	req := propertySearchRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrorResponse{Message: err.Error()})
	}

	if err := ctx.Validate(req); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			return ctx.JSON(common.ErrValidate.Code, common.ErrValidate)
		}

		return ctx.JSON(
			common.ErrValidate.Code, common.NewErrorValidateResponse(common.ErrValidate.Err, errors),
		)
	}

	properties, count, err := t.service.Search(ctx.Request().Context(), mapSearchRequestToSearch(&req))
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if len(properties) == 0 {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(
		http.StatusOK, common.DefaultResponse{
			Data: propertiesToListResponse(properties), Count: len(properties), TotalCount: count,
		},
	)
}

// SearchMarkers
//
// @Summary Search properties markers.
// @Description Search properties markers using different filters
// @Tags properties
// @Accept json
// @Produce json
// @Param Search body propertySearchRequest true "property search body"
// @Success 200 {object} common.DefaultResponse{data=[]propertyMarkerResponse}
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/properties/markers [POST]
func (t Handler) SearchMarkers(ctx echo.Context) error {
	req := propertySearchRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrorResponse{Message: err.Error()})
	}

	if err := ctx.Validate(req); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			return ctx.JSON(common.ErrValidate.Code, common.ErrValidate)
		}

		return ctx.JSON(
			common.ErrValidate.Code, common.NewErrorValidateResponse(common.ErrValidate.Err, errors),
		)
	}

	properties, count, err := t.service.Search(ctx.Request().Context(), mapSearchRequestToSearch(&req))
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if len(properties) == 0 {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(
		http.StatusOK, common.DefaultResponse{
			Data: propertiesToMarkerResponse(properties), Count: len(properties), TotalCount: count,
		},
	)
}
