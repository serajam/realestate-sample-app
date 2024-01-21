/* Copyright (C) Fedir Petryk */

package user_properties

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Activate
//
// @Summary Activate Property
// @Description Activate Property by ID
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
// @Router /v1/user/properties/{id}/activate  [PUT]
func (t Handler) Activate(ctx echo.Context) error {
	return t.processStatus(ctx, lo.ToPtr(true))
}

// Deactivate
//
// @Summary Deactivate Property
// @Description Deactivate Property by ID
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
// @Router /v1/user/properties/{id}/deactivate  [PUT]
func (t Handler) Deactivate(ctx echo.Context) error {
	return t.processStatus(ctx, lo.ToPtr(false))
}

func (t Handler) processStatus(ctx echo.Context, status *bool) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: err.Error()})
	}

	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: err.Error()})
	}

	prop, err := t.userPropService.GetUserProperty(ctx.Request().Context(), id, userID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if prop == nil {
		return ctx.JSON(http.StatusNotFound, common.EmptyResponse{})
	}

	prop.Active = status

	prop, err = t.userPropService.UpdateUserProperty(ctx.Request().Context(), prop)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}
