/* Copyright (C) Fedir Petryk */

package users

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// PutProfile
//
// @Summary PutProfile
// @Description PutProfile
// @Tags users
// @Accept json
// @Produce json
// @Success 204 {object} common.EmptyResponse
// @Param profile body profileRequest true "profile update"
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/user/profile  [PATCH]
func (h Handler) PutProfile(ctx echo.Context) error {
	req := profileRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrBind)
	}

	if err := ctx.Validate(req); err != nil {
		validErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return ctx.JSON(common.ErrValidate.Code, common.ErrValidate)
		}

		return ctx.JSON(
			common.ErrValidate.Code, common.NewErrorValidateResponse(common.ErrValidate.Err, validErrors),
		)
	}

	userID, err := common.UserID(ctx)
	if err != nil {
		return err
	}

	userDom := mapProfileRequestToUserDomain(&req)
	userDom.ID = userID

	err = h.profileService.UpdateProfile(ctx.Request().Context(), userDom)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}
