/* Copyright (C) Fedir Petryk */

package search_filters

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Create
//
// @Summary Create searchFilters
// @Description Create searchFilters
// @Tags searchfilters
// @Accept json
// @Produce json
// @Param searchFilters body searchFiltersRequest true "search filter data"
// @Success 200 {object} searchFiltersCreateResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/saved-searches  [POST]
func (t Handler) Create(ctx echo.Context) error {
	req := &searchFiltersRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrorResponse{Message: err.Error()})
	}

	if err := ctx.Validate(req); err != nil {
		errorsValidate, ok := err.(validator.ValidationErrors)
		if !ok {
			return ctx.JSON(common.ErrValidate.Code, common.ErrValidate)
		}

		return ctx.JSON(
			common.ErrValidate.Code, common.NewErrorValidateResponse(common.ErrValidate.Err, errorsValidate),
		)
	}

	userID, err := common.UserID(ctx)
	if err != nil {
		return err
	}

	searchFilter, err := mapSearchFiltersRequestToSearchFilters(req)
	if err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrorResponse{Message: err.Error()})
	}

	searchFilter.UserID = userID
	filter, err := t.service.Create(ctx.Request().Context(), searchFilter)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(
		http.StatusOK,
		searchFiltersCreateResponse{ID: filter.ID},
	)
}
