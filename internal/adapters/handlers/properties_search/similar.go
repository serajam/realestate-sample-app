/* Copyright (C) Fedir Petryk */

package properties_search

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Similar
//
// @Summary  List properties by similarity
// @Description  List properties by similarity base on specified property id.
// @Tags properties
// @Accept json
// @Produce json
// @Param Search body propertyListRequest true "property list body"
// @Success 200 {object} common.DefaultResponse{data=[]propertyResponse}
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/properties/{id}/similar [POST]
func (t Handler) Similar(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: err.Error()})
	}

	req := propertyGetRequest{ID: id}
	if err := ctx.Validate(req); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			return ctx.JSON(common.ErrValidate.Code, common.ErrValidate)
		}

		return ctx.JSON(
			common.ErrValidate.Code, common.NewErrorValidateResponse(common.ErrValidate.Err, errors),
		)
	}

	properties, err := t.service.SearchSimilar(ctx.Request().Context(), id)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if len(properties) == 0 {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(
		http.StatusOK, common.DefaultResponse{Data: propertiesToListResponse(properties), Count: len(properties)},
	)
}
