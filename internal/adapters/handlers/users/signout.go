/* Copyright (C) Fedir Petryk */

package users

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// SignOut
//
// @Summary SignOut User
// @Description SignOut User
// @Tags auth
// @Accept json
// @Produce json
// @Success 204 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/sign-out  [PUT]
func (h Handler) SignOut(ctx echo.Context) error {
	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: "user_id is empty"})
	}

	deviceID := ctx.Request().Header.Get("X-Device-Id")
	err := h.authService.SignOut(ctx.Request().Context(), userID, deviceID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}

// SignOutAll
//
// @Summary SignOut User
// @Description SignOut User
// @Tags auth
// @Accept json
// @Produce json
// @Success 204 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/sign-out/all  [PUT]
func (h Handler) SignOutAll(ctx echo.Context) error {
	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: "user_id is empty"})
	}

	err := h.authService.SignOutAll(ctx.Request().Context(), userID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}
