package service

import (
	"context"
	"time"
	"github.com/telcoflow/telcoflow/services/partner-service/internal/domain"
)

type PartnerService struct{}

func NewPartnerService() *PartnerService {
	return &PartnerService{}
}

func (s *PartnerService) OnboardPartner(ctx context.Context, p *domain.Partner) error {
	p.ID = "PRT-" + time.Now().Format("20060102")
	p.OnboardDate = time.Now()
	p.Status = "Active"
	return nil
}
