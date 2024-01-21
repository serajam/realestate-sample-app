/* Copyright (C) Fedir Petryk */

package common

import (
	"github.com/labstack/echo/v4"
)

const (
	ID         = "id"
	UserId     = "user_id"
	PropertyId = "property_id"
	ImageId    = "image_id"
)

func UserID(ctx echo.Context) (int, error) {
	idStr := ctx.Get(UserId)
	userID := idStr.(int)
	if userID == 0 {
		return 0, ctx.JSON(ErrValidate.Code, ErrorResponse{Message: "INVALID_USER_ID"})
	}

	return userID, nil
}

func IDFromPath(c echo.Context, key string) (int, error) {
	var id int
	if err := echo.PathParamsBinder(c).Int(key, &id).BindError(); err != nil {
		return 0, c.JSON(ErrValidate.Code, ErrorResponse{Message: "INVALID_PATH_ID"})
	}

	if id == 0 {
		return 0, c.JSON(ErrValidate.Code, ErrorResponse{Message: "EMPTY_PATH_ID"})
	}

	return id, nil
}

func IDFromContext(c echo.Context, key string) (int, error) {
	idStr := c.Get(key)
	id := idStr.(int)

	if id == 0 {
		return 0, c.JSON(ErrValidate.Code, ErrorResponse{Message: "INVALID_ID"})
	}

	return id, nil
}
