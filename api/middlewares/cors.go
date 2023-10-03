package middlewares

import (
	"clean-go-echo/library"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CorsMiddleware middleware for cors
type CorsMiddleware struct {
	handler library.RequestHandler
	logger  library.LoggerZap
	env     library.Env
}

// ModuleCorsMiddleware creates new cors middleware
func ModuleCorsMiddleware(handler library.RequestHandler, logger library.LoggerZap, env library.Env) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

// Setup sets up cors middleware
func (m CorsMiddleware) Setup() {

	m.handler.Echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			req := c.Request()
			res := c.Response()

			allowList := map[string]bool{
				"http://localhost:3000": true,
				"":                      true,
			}

			log.Println(req.Header.Get("Origin"))

			if origin := req.Header.Get("Origin"); allowList[origin] {
				res.Header().Add("Access-Control-Allow-Origin", origin)
				res.Header().Add("Access-Control-Allow-Methods", "*")
				res.Header().Add("Access-Control-Allow-Headers", "*")
				res.Header().Add("Content-Type", "application/json")

				if req.Method != "OPTIONS" {
					err := next(c)
					if err != nil {
						c.Error(err)
					}
				}
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Cors Origin")
			}

			return nil

		}
	})
}
