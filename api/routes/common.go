package routes

import (
	"clean-go-echo/api/controllers"
	"clean-go-echo/library"

	"github.com/labstack/echo/v4"
)

// struct routes
type CommonRoutes struct {
	handler        library.RequestHandler
	userController controllers.UserController
}

// Setup routes
func (s CommonRoutes) Setup() {
	api := s.handler.Echo.Group("/")
	{
		api.GET("health-check", func(c echo.Context) error {
			return c.JSON(200, "OK")
		})
		api.GET("user", s.userController.GetUser)
		api.GET("users", s.userController.GetUserSecond)
	}
}

func ModuleCommonRoutes(
	handler library.RequestHandler,
	userConteroller controllers.UserController,
) CommonRoutes {
	return CommonRoutes{
		handler:        handler,
		userController: userConteroller,
	}
}
