package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Role string

const (
	RoleAdmin      Role = "admin"
	RoleCSR        Role = "csr"
	RoleNetworkEng Role = "network_eng"
	RoleBillingMgr Role = "billing_mgr"
	RoleCustomer   Role = "customer"
)

type Permission string

const (
	PermOrderCreate    Permission = "order:create"
	PermOrderRead      Permission = "order:read"
	PermCustomerUpdate Permission = "customer:update"
	PermBillingView    Permission = "billing:view"
)

// RBACManager handles role-based access control logic
type RBACManager struct {
	rolePermissions map[Role][]Permission
}

func NewRBACManager() *RBACManager {
	return &RBACManager{
		rolePermissions: map[Role][]Permission{
			RoleAdmin:      {PermOrderCreate, PermOrderRead, PermCustomerUpdate, PermBillingView},
			RoleCSR:        {PermOrderRead, PermCustomerUpdate},
			RoleNetworkEng: {PermOrderRead},
			RoleBillingMgr: {PermBillingView},
			RoleCustomer:   {PermOrderRead},
		},
	}
}

func (m *RBACManager) HasPermission(role Role, perm Permission) bool {
	permissions, ok := m.rolePermissions[role]
	if !ok {
		return false
	}
	for _, p := range permissions {
		if p == perm {
			return true
		}
	}
	return false
}

// RBACMiddleware checks if the user in the context has the required permission
func RBACMiddleware(manager *RBACManager, requiredPermission Permission) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// In a real system, the role would be extracted from the JWT claims
			// stored in the context by an authentication middleware.
			role, ok := c.Get("role").(Role)
			if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error":   "unauthorized",
				"message": "Missing role in context",
			})
			}

			if !manager.HasPermission(role, requiredPermission) {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error":   "insufficient permissions",
					"message": "Required permission: " + string(requiredPermission),
				})
			}

			return next(c)
		}
	}
}
