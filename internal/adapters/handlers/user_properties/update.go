/* Copyright (C) Fedir Petryk */

package user_properties

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Update
//
// @Summary Update PropertyUpdateRequest
// @Description Update PropertyUpdateRequest by ID
// @Tags properties
// @Accept json
// @Produce json
// @Param updateProperty body propertyCreateRequest true "property data"
// @Success 200 {object} common.DefaultResponse{data=propertyResponse}
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/user/properties/{id}  [PUT]
func (t Handler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: err.Error()})
	}

	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: err.Error()})
	}

	req := &propertyUpdateRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(common.ErrBind.Code, common.ErrorResponse{Message: err.Error()})
	}

	ok, err := t.ValidateRequest(ctx, req)
	if err != nil || !ok {
		return err
	}

	property, err := updatePropertyToDomain(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
	}
	property.ID = id
	property.UserID = userID

	property, err = t.userPropService.UpdateUserProperty(ctx.Request().Context(), property)
	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if property == nil {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(http.StatusOK, common.DefaultResponse{Data: propertyToResponse(property)})
}
