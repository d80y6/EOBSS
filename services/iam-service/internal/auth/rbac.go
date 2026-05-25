package iamauth

import (
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/auth"
)

// Re-exporting from go-common to maintain compatibility with existing code
// while moving the source of truth to the shared library.

type Role = auth.Role

const (
	RoleAdmin      = auth.RoleAdmin
	RoleCSR        = auth.RoleCSR
	RoleNetworkEng = auth.RoleNetworkEng
	RoleBillingMgr = auth.RoleBillingMgr
	RoleCustomer   = auth.RoleCustomer
)

type Permission = auth.Permission

const (
	PermOrderCreate    = auth.PermOrderCreate
	PermOrderRead      = auth.PermOrderRead
	PermCustomerUpdate = auth.PermCustomerUpdate
	PermBillingView    = auth.PermBillingView
)

type RBACManager = auth.RBACManager

func NewRBACManager() *RBACManager {
	return auth.NewRBACManager()
}
