package routes

import (
	"clean-go-echo/api/controllers"
	"clean-go-echo/api/middlewares"
	"clean-go-echo/library"
)

// UserRoutes struct routes
type UserRoutes struct {
	handler        library.RequestHandler
	userController controllers.UserController
	middleware     middlewares.JWTAuthMiddleware
}

// Setup routes
func (s UserRoutes) Setup() {
	api := s.handler.Echo.Group("user/")
	{
		api.POST("store", s.userController.StoreUser)
		api.GET("list", s.userController.ListUser)
	}
}

// ModuleCommonRoutes
func ModuleUserRoutes(
	handler library.RequestHandler,
	userConteroller controllers.UserController,
	middleware middlewares.JWTAuthMiddleware,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userConteroller,
		middleware:     middleware,
	}
}
