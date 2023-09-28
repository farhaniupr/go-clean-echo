package commands

import (
	"clean-go-echo/api/controllers"
	"clean-go-echo/api/routes"
	"clean-go-echo/library"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	library.Module,
	routes.Module,

// 	routes.Module,
// 	lib.Module,
// 	services.Module,
// 	middlewares.Module,
// 	repository.Module,
)
