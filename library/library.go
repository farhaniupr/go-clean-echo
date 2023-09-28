package library

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(ModuleEcho),
	fx.Provide(ModuleDatabase),
	fx.Provide(ModuleEnv),
)
