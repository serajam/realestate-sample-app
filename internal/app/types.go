/* Copyright (C) Fedir Petryk */

package app

import "github.com/labstack/echo/v4"

type Subscriber interface {
	Shutdown()
}

type Handler interface {
	AddRoutes(e *echo.Group, m ...echo.MiddlewareFunc)
}
