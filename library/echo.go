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
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, error=${error}, status=${status}\n",
	}))
	engine.Use(middleware.Recover())
	return RequestHandler{Echo: engine}
}
