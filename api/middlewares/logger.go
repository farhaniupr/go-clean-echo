package middlewares

import (
	"clean-go-echo/library"

	"github.com/labstack/echo/v4"
)

// DatabaseTrx middleware for transactions support for database
type LoggerMiddlewre struct {
	handler library.RequestHandler
	env     library.Env
}

// ModuleLogger creates new database transactions middleware
func ModuleLogger(
	handler library.RequestHandler,
	env library.Env,
) LoggerMiddlewre {
	return LoggerMiddlewre{
		handler: handler,
		env:     env,
	}
}

// Setup sets up logger middleware
func (m LoggerMiddlewre) Setup() {

	m.handler.Echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			res := c.Response()
			n := res.Status

			switch {
			case n >= 500:
				library.Writelog(c, m.env, "err", "server error")
			case n >= 400:
				library.Writelog(c, m.env, "warn", "client error")
			default:
				library.Writelog(c, m.env, "info", "success")
			}

			return nil

		}
	})
}
