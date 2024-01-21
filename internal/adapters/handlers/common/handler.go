/* Copyright (C) Fedir Petryk */

package common

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	AddRoutes(e *echo.Group, m ...echo.MiddlewareFunc)
	ValidateRequest(ctx echo.Context, req interface{}) (bool, error)
}

type BaseHandler struct {
}

func (n BaseHandler) ValidateRequest(ctx echo.Context, req interface{}) (bool, error) {
	if err := ctx.Validate(req); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, ctx.JSON(ErrValidate.Code, ErrValidate)
		}

		return false, ctx.JSON(
			ErrValidate.Code, NewErrorValidateResponse(ErrValidate.Err, errors),
		)
	}

	return true, nil
}
