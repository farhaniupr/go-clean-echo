package routes

import "go.uber.org/fx"

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(ModuleCommonRoutes),
	fx.Provide(ModuleRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func ModuleRoutes(
	commonRoutes CommonRoutes,
) Routes {
	return Routes{
		commonRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
