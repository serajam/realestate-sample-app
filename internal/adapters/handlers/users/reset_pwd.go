/* Copyright (C) Fedir Petryk */

package users

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// ResetPwd
//
// @Summary Reset Password
// @Description User password reset confirmation request
// @Tags users
// @Accept json
// @Param resetEmail body resetPwdRequest true "user email"
// @Success 204 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/reset-password  [POST]
func (h Handler) ResetPwd(ctx echo.Context) error {
	req := resetPwdRequest{}
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

	err := h.pwdResetService.ResetPassword(ctx.Request().Context(), req.Email)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}

// NewPwd
//
// @Summary New Password
// @Description Set password reset confirmation request
// @Tags users
// @Accept json
// @Param setPwd body newPasswordRequest true "user pwd"
// @Param token path string true "pwd reset token"
// @Success 204 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/set-password/{token}  [PUT]
func (h Handler) NewPwd(ctx echo.Context) error {
	req := newPasswordRequest{}
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

	err := h.pwdResetService.SetNewPassword(ctx.Request().Context(), req.Token, req.Password)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}
