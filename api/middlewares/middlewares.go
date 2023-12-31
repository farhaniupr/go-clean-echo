package middlewares

import (
	"go.uber.org/fx"
)

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(ModuleDatabase),
	fx.Provide(ModuleMiddlewares),
	fx.Provide(ModuleJWTAuthMiddleware),
	fx.Provide(ModuleLogger),
	fx.Provide(ModuleCorsMiddleware),
)

// IMiddleware middleware interface
type IMiddleware interface {
	Setup()
}

// Middlewares contains multiple middleware
type Middlewares []IMiddleware

// ModuleMiddlewares creates new middlewares
// Register the middleware that should be applied directly (globally)
func ModuleMiddlewares(
	jwtauthMiddleware JWTAuthMiddleware,
	loggerMiddleware LoggerMiddlewre,
	corsMiddleware CorsMiddleware,
) Middlewares {
	return Middlewares{
		jwtauthMiddleware,
		loggerMiddleware,
		corsMiddleware,
	}
}

// Setup sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
