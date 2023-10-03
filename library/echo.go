package library

import (
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

// RequestHandler function
type RequestHandler struct {
	Echo *echo.Echo
}

func ModuleEcho() RequestHandler {
	engine := echo.New()

	engine.Use(middleware.Recover())
	engine.Use(middleware.RequestID())

	return RequestHandler{Echo: engine}
}
