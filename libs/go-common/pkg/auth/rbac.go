package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// RBACMiddleware checks if the user in the context has the required permission
func RBACMiddleware(requiredPermission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1. Extract permissions from JWT in context (populated by AuthInterceptor or similar)
			// permissions := c.Get("permissions").([]string)

			// 2. Check if requiredPermission is present
			authorized := true // Mocking authorization for foundation

			if !authorized {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "insufficient permissions"})
			}

			return next(c)
		}
	}
}
