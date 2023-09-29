package routes

import (
	"clean-go-echo/library"

	"github.com/labstack/echo/v4"
)

// CommonRoutes struct routes
type CommonRoutes struct {
	handler library.RequestHandler
	// middleware middlewares.JWTAuthMiddleware
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

// ModuleCommonRoutes
func ModuleCommonRoutes(
	handler library.RequestHandler,
	// middleware middlewares.JWTAuthMiddleware,
) CommonRoutes {
	return CommonRoutes{
		handler: handler,
		// middleware: middleware,
	}
}
