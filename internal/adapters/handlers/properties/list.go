/* Copyright (C) Fedir Petryk */

package properties

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// List
//
// @Summary List properties by ids.
// @Description List properties by ids.
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
// @Router /v1/properties/list [POST]
func (t Handler) List(ctx echo.Context) error {
	req := propertyListRequest{}
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

	properties, err := t.propService.List(ctx.Request().Context(), req.IDs)
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
