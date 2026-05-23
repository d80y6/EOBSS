package tenant

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type contextKey string

const TenantIDKey contextKey = "tenant_id"

// TenantMiddleware extracts the tenant ID from the X-Tenant-ID header
func TenantMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tenantID := c.Request().Header.Get("X-Tenant-ID")
		if tenantID == "" {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "tenant context missing"})
		}

		ctx := context.WithValue(c.Request().Context(), TenantIDKey, tenantID)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func FromContext(ctx context.Context) string {
	if id, ok := ctx.Value(TenantIDKey).(string); ok {
		return id
	}
	return ""
}
