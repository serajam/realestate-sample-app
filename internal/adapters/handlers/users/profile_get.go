/* Copyright (C) Fedir Petryk */

package users

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// GetProfile
//
// @Summary GetProfile
// @Description GetProfile
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} profileResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/user/profile  [GET]
func (h Handler) GetProfile(ctx echo.Context) error {
	userID, err := common.UserID(ctx)
	if err != nil {
		return err
	}

	profile, err := h.profileService.GetProfile(ctx.Request().Context(), userID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, mapUserDomainToProfile(profile))
}
