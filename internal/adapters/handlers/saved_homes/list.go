/* Copyright (C) Fedir Petryk */

package saved_homes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// List
//
// @Summary List saved properties
// @Description List properties by ids.
// @Tags savedhomes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Pagination body paginationRequest true "pagination"
// @Success 200 {object} common.DefaultResponse{data=[]propertyResponse}
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/saved-homes/list [POST]
func (t Handler) List(ctx echo.Context) error {
	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid user searchID"})
	}

	req := paginationRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrorResponse{Message: err.Error()})
	}

	ok, err := t.ValidateRequest(ctx, req)
	if err != nil || !ok {
		return err
	}

	properties, err := t.service.List(ctx.Request().Context(), userID, mapPaginationRequestToPagination(req))
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if len(properties) == 0 {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(
		http.StatusOK, common.DefaultResponse{Data: mapPropertiesToListResponse(properties), Count: len(properties)},
	)
}
