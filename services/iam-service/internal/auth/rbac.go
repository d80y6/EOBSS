package auth

import (
	"context"
)

type Role string

const (
	RoleAdmin       Role = "admin"
	RoleCSR         Role = "csr"         // Customer Service Representative
	RoleNetworkEng  Role = "network_eng" // Network Engineer
	RoleBillingMgr  Role = "billing_mgr"
	RoleCustomer    Role = "customer"
)

type Permission string

const (
	PermOrderCreate    Permission = "order:create"
	PermOrderRead      Permission = "order:read"
	PermCustomerUpdate Permission = "customer:update"
	PermBillingView    Permission = "billing:view"
)

// RBACManager handles role-based access control
type RBACManager struct {
	rolePermissions map[Role][]Permission
}

func NewRBACManager() *RBACManager {
	return &RBACManager{
		rolePermissions: map[Role][]Permission{
			RoleAdmin: {PermOrderCreate, PermOrderRead, PermCustomerUpdate, PermBillingView},
			RoleCSR:   {PermOrderRead, PermCustomerUpdate},
			RoleCustomer: {PermOrderRead},
		},
	}
}

func (m *RBACManager) HasPermission(ctx context.Context, role Role, perm Permission) bool {
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
