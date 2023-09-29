package middlewares

import (
	"clean-go-echo/services"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	// "github.com/gin-gonic/gin"
)

// JWTAuthMiddleware middleware for jwt authentication
type JWTAuthMiddleware struct {
	service services.JWTAuth_MethodService
	// logger  lib.Logger
}

// ModuleJWTAuthMiddleware creates new jwt auth middleware
func ModuleJWTAuthMiddleware(
	// logger lib.Logger,
	service services.JWTAuth_MethodService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		service: service,
		// logger:  logger,
	}
}

// Setup sets up jwt auth middleware
func (m JWTAuthMiddleware) Setup() {}

// Handler handles middleware functionality
func (m JWTAuthMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			user, authorized, err := m.service.Authorize(authToken)
			_ = user
			if authorized {
				next(c)
				return nil
			}
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})

			return nil
		}
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "you are not authorized",
		})
		return nil
	}
}
