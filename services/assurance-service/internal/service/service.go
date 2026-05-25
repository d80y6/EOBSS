package service

import (
	"context"
	"github.com/telcoflow/telcoflow/services/assurance-service/internal/correlation"
	"github.com/telcoflow/telcoflow/services/assurance-service/internal/domain"
)

type AssuranceService struct {
	correlationEngine *correlation.AlarmCorrelationEngine
}

func NewAssuranceService(engine *correlation.AlarmCorrelationEngine) *AssuranceService {
	return &AssuranceService{correlationEngine: engine}
}

func (s *AssuranceService) ProcessAlarm(ctx context.Context, alarm *domain.Alarm) error {
	// Standard TMF642 Alarm processing
	return s.correlationEngine.Correlate(ctx, alarm)
}
