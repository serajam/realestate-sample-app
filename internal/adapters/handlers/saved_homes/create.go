/* Copyright (C) Fedir Petryk */

package saved_homes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

// Create
//
// @Summary Add saved home
// @Description Adds property to user saved homes
// @Accept json
// @Tags savedhomes
// @Produce json
// @Success 200 {object} propertyResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 403 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security ApiKeyAuth
// @Router /v1/saved-homes/{id}  [PUT]
func (t Handler) Create(ctx echo.Context) error {
	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid user searchID"})
	}

	propId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(common.ErrValidate.Code, common.ErrorResponse{Message: err.Error()})
	}

	prop, err := t.service.Add(
		ctx.Request().Context(), properties.UserSavedHome{
			UserID:     userID,
			PropertyID: propId,
		},
	)

	if err != nil {
		return common.ProcessDomainError(ctx, err)
	}

	if prop == nil {
		return ctx.JSON(http.StatusNoContent, common.EmptyResponse{})
	}

	return ctx.JSON(http.StatusOK, mapPropertyToPropertyResponse(prop))
}
