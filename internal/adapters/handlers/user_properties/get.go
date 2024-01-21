/* Copyright (C) Fedir Petryk */

package user_properties

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// GetUserProp user prop
//
// @Summary GetUserProp user propertyGetRequest by ID
// @Description GetUserProp user propertyGetRequest by ID
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
// @Router /v1/user/properties/{id}  [GET]
func (t Handler) GetUserProp(ctx echo.Context) error {
	userID, err := common.UserID(ctx)
	if err != nil {
		return err
	}

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

	prop, err := t.userPropService.GetUserProperty(ctx.Request().Context(), req.ID, userID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if prop == nil {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(http.StatusOK, common.DefaultResponse{Data: propertyToResponse(prop)})
}
