/* Copyright (C) Fedir Petryk */

package users

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Deactivate
//
// @Summary Deactivate Profile
// @Description Deactivate Profile
// @Tags users
// @Accept json
// @Produce json
// @Success 204 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/user/profile/deactivate  [POST]
func (h Handler) Deactivate(ctx echo.Context) error {
	userID, err := common.UserID(ctx)
	if err != nil {
		return err
	}

	err = h.profileService.Deactivate(ctx.Request().Context(), userID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}
