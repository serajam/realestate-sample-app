/* Copyright (C) Fedir Petryk */

package common

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	domainErrors "github.com/serajam/realestate-sample-app/internal/core/domain/errors"
)

func ErrorHandler(logger *zap.SugaredLogger) func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		// try to find code error type
		var codeErr Error
		if !errors.As(err, &codeErr) {
			// try echo HTTPError type
			var he *echo.HTTPError
			if errors.As(err, &he) {
				codeErr.Code = he.Code
				if m, ok := he.Message.(string); ok {
					codeErr.Err = m
				}
			} else {
				// response shouldn't show unexpected errors, so display 500
				codeErr = ErrInternal
			}
		}

		// SignUp response
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(codeErr.Code)
		} else {
			err = c.JSON(codeErr.Code, ErrorResponse{Message: codeErr.Err})
		}
		if err != nil {
			logger.Error("write error response", "error", err)
		}
	}
}

func ProcessDomainError(ctx echo.Context, err error) error {
	switch {
	case errors.As(err, &domainErrors.RequestTimeout):
		return ctx.JSON(ErrTimeout.Code, ErrTimeout)
	case errors.As(err, &domainErrors.NotFound):
		return ctx.JSON(ErrNotFound.Code, ErrNotFound)
	case errors.As(err, &domainErrors.AccessDenied):
		return ctx.JSON(ErrAccessForbidden.Code, ErrAccessForbidden)
	case errors.As(err, &domainErrors.Exists):
		return ctx.JSON(ErrExists.Code, ErrExists)
	case errors.As(err, &domainErrors.OpFail{}):
		return ctx.JSON(ErrInternal.Code, ErrInternal)
	case errors.As(err, &domainErrors.Internal):
		return ctx.JSON(ErrInternal.Code, ErrInternal)
	case errors.As(err, &domainErrors.Dummy):
		return ctx.JSON(http.StatusNoContent, EmptyResponse{})
	case errors.As(err, &domainErrors.User{}):
		return ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{Message: err.Error()})
	default:
		return ctx.JSON(ErrInternal.Code, ErrInternal)
	}
}
