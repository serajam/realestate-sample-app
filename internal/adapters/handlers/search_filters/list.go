/* Copyright (C) Fedir Petryk */

package search_filters

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// List
//
// @Summary List users search filters.
// @Description List users search filters.
// @Tags searchfilters
// @Accept json
// @Produce json
// @Success 200 {object} common.DefaultResponse{data=[]searchFiltersResponse}
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/saved-searches [GET]
func (t Handler) List(ctx echo.Context) error {
	userID, err := common.UserID(ctx)
	if err != nil {
		return err
	}

	savedSearchFilters, err := t.service.List(ctx.Request().Context(), userID)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if len(savedSearchFilters) == 0 {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(
		http.StatusOK,
		common.DefaultResponse{
			Data: mapSearchFiltersToSearchFiltersResponse(savedSearchFilters), Count: len(savedSearchFilters),
		},
	)
}
