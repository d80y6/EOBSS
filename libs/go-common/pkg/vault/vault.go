package vault

import (
	"context"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"go.uber.org/zap"
)

// VaultClient provides a carrier-grade interface for secret management
type VaultClient struct {
	client *vault.Client
}

func NewVaultClient(address, token string) (*VaultClient, error) {
	config := vault.DefaultConfig()
	config.Address = address

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, err
	}
	client.SetToken(token)

	return &VaultClient{client: client}, nil
}

// GetSecret retrieves a secret from a specific path
func (v *VaultClient) GetSecret(ctx context.Context, path string) (map[string]interface{}, error) {
	logger.Info("Retrieving secret from Vault", zap.String("path", path))

	secret, err := v.client.Logical().Read(path)
	if err != nil {
		return nil, err
	}
	if secret == nil {
		return nil, fmt.Errorf("secret not found at path: %s", path)
	}

	return secret.Data, nil
}
