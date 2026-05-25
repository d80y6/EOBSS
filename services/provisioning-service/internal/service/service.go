package service

import (
	"context"
	"github.com/telcoflow/telcoflow/services/provisioning-service/internal/adapter"
)

type ProvisioningService struct {
	radiusAdapter  *adapter.RadiusAdapter
	kamailioAdapter *adapter.KamailioAdapter
	open5gsAdapter  *adapter.Open5GSAdapter
}

func NewProvisioningService(ra *adapter.RadiusAdapter, ka *adapter.KamailioAdapter, oa *adapter.Open5GSAdapter) *ProvisioningService {
	return &ProvisioningService{
		radiusAdapter:  ra,
		kamailioAdapter: ka,
		open5gsAdapter:  oa,
	}
}

func (s *ProvisioningService) ProvisionRadius(ctx context.Context, username, password string) error {
	return s.radiusAdapter.ProvisionUser(ctx, username, password, nil)
}

func (s *ProvisioningService) ProvisionVoIP(ctx context.Context, username, domain, password string) error {
	return s.kamailioAdapter.ProvisionSIPUser(ctx, username, domain, password)
}

func (s *ProvisioningService) Provision5G(ctx context.Context, imsi, msisdn string) error {
	return s.open5gsAdapter.Provision5GSubscriber(ctx, imsi, msisdn)
}
