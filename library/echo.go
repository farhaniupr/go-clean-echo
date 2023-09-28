package library

import (
	"github.com/labstack/echo/v4"
)

// RequestHandler function
type RequestHandler struct {
	Echo *echo.Echo
}

func ModulEcho() RequestHandler {
	engine := echo.New()
	return RequestHandler{Echo: engine}
}
