/* Copyright (C) Fedir Petryk */

package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

func Auth(a AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenStr := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
			if tokenStr == "" {
				return common.ErrUnauthorized
			}

			// deviceId := c.Request().Header.Get("Device-Id")
			// if deviceId == "" {
			// 	deviceId = "default"
			// 	//	return c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Device-Id header is required"})
			// }

			token, err := a.Validate(c.Request().Context(), tokenStr)
			if err != nil {
				return common.ErrUnauthorized
			}

			c.Set("user_id", token.UserID)
			c.Set("token_uuid", token.TokenUUID)

			return next(c)
		}
	}
}
