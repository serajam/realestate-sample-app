/* Copyright (C) Fedir Petryk */

package user_properties

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// ListByUser
//
// @Summary List properties by auth user id.
// @Description List properties auth user id.
// @Tags properties
// @Accept json
// @Produce json
// @Param Search body propertyListRequest true "property list body"
// @Success 200 {object} common.DefaultResponse{data=[]propertyResponse}
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/user/properties [GET]
func (t Handler) ListByUser(ctx echo.Context) error {
	userID, err := common.UserID(ctx)
	if err != nil {
		return err
	}

	req := common.BaseSearchRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrorResponse{Message: err.Error()})
	}

	ok, err := t.ValidateRequest(ctx, req)
	if err != nil || !ok {
		return err
	}

	reqDom := listUserPropertiesToListRequest(req)
	properties, err := t.userPropService.ListByUser(ctx.Request().Context(), userID, reqDom)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if len(properties) == 0 {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(
		http.StatusOK, common.DefaultResponse{Data: propertiesToListResponse(properties), Count: len(properties)},
	)
}
