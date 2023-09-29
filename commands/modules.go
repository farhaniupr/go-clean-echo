package commands

import (
	"clean-go-echo/api/controllers"
	"clean-go-echo/api/middlewares"
	"clean-go-echo/api/routes"
	"clean-go-echo/library"
	"clean-go-echo/repository"
	"clean-go-echo/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	library.Module,
	controllers.Module,
	routes.Module,
	services.Module,
	repository.Module,
	middlewares.Module,
)
