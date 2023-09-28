package routes

import (
	"clean-go-echo/library"

	"github.com/labstack/echo/v4"
)

// struct routes
type CommonRoutes struct {
	handler library.RequestHandler
}

// Setup routes
func (s CommonRoutes) Setup() {
	api := s.handler.Echo.Group("/")
	{
		api.GET("health-check", func(c echo.Context) error {
			return c.JSON(200, "OK")
		})
	}
}

func ModuleCommonRoutes(
	handler library.RequestHandler,
) CommonRoutes {
	return CommonRoutes{
		handler: handler,
	}
}
