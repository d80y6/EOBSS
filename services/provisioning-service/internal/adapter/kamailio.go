package adapter

import (
	"context"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"go.uber.org/zap"
)

// KamailioAdapter handles VoIP subscriber provisioning for SIP servers
type KamailioAdapter struct {
	DBConn string
}

func (a *KamailioAdapter) ProvisionSIPUser(ctx context.Context, username, domain, password string) error {
	logger.Info("Provisioning SIP user in Kamailio",
		zap.String("username", username),
		zap.String("domain", domain),
	)

	// Implementation would perform an SQL INSERT into 'subscriber' table
	// or call Kamailio RPC/JSON API.
	return nil
}
