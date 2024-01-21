/* Copyright (C) Fedir Petryk */

package properties

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Get
//
// @Summary Get propertyGetRequest by ID
// @Description Get propertyGetRequest by ID
// @Tags properties
// @Accept json
// @Produce json
// @Param id path int true "propertyGetRequest ID"
// @Success 200 {object} common.DefaultResponse{data=propertyResponse}
// @Success 404 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/properties/{id}  [GET]
func (t Handler) Get(ctx echo.Context) error {
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

	prop, err := t.propService.Get(ctx.Request().Context(), req.ID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if prop == nil {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	user, err := t.profileService.GetProfile(ctx.Request().Context(), prop.UserID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, common.DefaultResponse{Data: propertyWithBroker(propertyToResponse(prop), user)})
}
