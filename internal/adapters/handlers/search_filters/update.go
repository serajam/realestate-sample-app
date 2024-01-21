/* Copyright (C) Fedir Petryk */

package search_filters

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Update
//
// @Summary Update UserSearchFilters
// @Description Update UserSearchFilters by ID
// @Tags searchfilters
// @Accept json
// @Produce json
// @Param searchFilters body searchFiltersRequest true "search filters data"
// @Param id path int true "UserSearchFilters filters ID"
// @Success 204 {object} common.EmptyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/saved-searches/{id}  [PUT]
func (t Handler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: err.Error()})
	}

	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: "INVALID_USER_ID"})
	}

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

	searchFilters, err := mapSearchFiltersRequestToSearchFilters(req)
	if err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrorResponse{Message: err.Error()})
	}

	searchFilters.ID = id
	searchFilters.UserID = userID

	err = t.service.Update(ctx.Request().Context(), searchFilters)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
}
