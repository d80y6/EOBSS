package adapter

import (
	"context"
	"fmt"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"go.uber.org/zap"
)

// RadiusAdapter handles communication with Radius servers (e.g. FreeRADIUS)
type RadiusAdapter struct {
	Endpoint string
	Secret   string
}

func NewRadiusAdapter(endpoint, secret string) *RadiusAdapter {
	return &RadiusAdapter{
		Endpoint: endpoint,
		Secret:   secret,
	}
}

// ProvisionUser creates an authentication record in the Radius server
func (a *RadiusAdapter) ProvisionUser(ctx context.Context, username, password string, profile map[string]string) error {
	logger.Info("Provisioning user in Radius",
		zap.String("username", username),
		zap.String("endpoint", a.Endpoint),
	)

	// In a real implementation, this would use a Radius client library
	// or perform a REST/SQL call if using a modern Radius backend.
	fmt.Printf("RADIUS: radclient -f user_config %s auth %s\n", a.Endpoint, a.Secret)

	return nil
}

// DeprovisionUser removes or disables the user in Radius
func (a *RadiusAdapter) DeprovisionUser(ctx context.Context, username string) error {
	logger.Info("Deprovisioning user from Radius", zap.String("username", username))
	return nil
}
