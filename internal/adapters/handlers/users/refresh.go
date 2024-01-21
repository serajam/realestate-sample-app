/* Copyright (C) Fedir Petryk */

package users

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Refresh
//
// @Summary Refresh User Access and Refresh token
// @Description Refresh User Access Refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} authToken
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/refresh-token  [PUT]
func (h Handler) Refresh(ctx echo.Context) error {
	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: "user_id is empty"})
	}

	tokenUUID := ctx.Get("token_uuid").(string)
	deviceID := ctx.Request().Header.Get("X-Device-Id")

	accessToken, refreshToken, err := h.authService.Refresh(ctx.Request().Context(), userID, tokenUUID, deviceID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, authToken{AccessToken: string(accessToken), RefreshToken: string(refreshToken)})
}
