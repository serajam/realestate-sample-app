/* Copyright (C) Fedir Petryk */

package users

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// SignIn
//
// @Summary SignIn User
// @Description SignIn User
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} authToken
// @Param signInUser body signIn true "user login"
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/sign-in  [POST]
func (h Handler) SignIn(ctx echo.Context) error {
	req := signIn{}
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

	deviceID := ctx.Request().Header.Get("X-Device-Id")

	accessToken, refreshToken, err := h.authService.Authenticate(
		ctx.Request().Context(), req.Email, req.Password, deviceID,
	)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, authToken{AccessToken: string(accessToken), RefreshToken: string(refreshToken)})
}
