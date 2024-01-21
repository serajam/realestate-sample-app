/* Copyright (C) Fedir Petryk */

package search_filters

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Delete
//
// @Summary Delete user search filter.
// @Description Delete user search filter.
// @Tags searchfilters
// @Accept json
// @Produce json
// @Param id path int true "Search filters ID"
// @Success 204 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/saved-searches/{id} [DELETE]
func (t Handler) Delete(ctx echo.Context) error {
	userID, err := common.UserID(ctx)
	if err != nil {
		return err
	}
	searchID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	err = t.service.Delete(ctx.Request().Context(), searchID, userID)
	if err != nil {
		return ctx.JSON(common.ErrInternal.Code, common.ErrInternal)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}
