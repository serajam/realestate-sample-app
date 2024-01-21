/* Copyright (C) Fedir Petryk */

package user_properties

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Delete
//
// @Summary Delete Property
// @Description Delete Property and images
// @Tags properties
// @Accept json
// @Produce json
// @Success 204 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/user/properties/{id}  [DELETE]
func (t Handler) Delete(ctx echo.Context) error {
	id, err := common.IDFromPath(ctx, common.ID)
	if err != nil {
		return err
	}

	userID, err := common.IDFromContext(ctx, common.UserId)
	if err != nil {
		return err
	}

	property, err := t.userPropService.GetUserProperty(ctx.Request().Context(), id, userID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	err = t.imgService.DeleteAll(ctx.Request().Context(), property)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	err = t.userPropService.DeleteUserProperty(ctx.Request().Context(), property)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}
