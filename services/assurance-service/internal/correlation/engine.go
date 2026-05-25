package correlation

import (
	"context"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/services/assurance-service/internal/domain"
	"go.uber.org/zap"
)

// AlarmCorrelationEngine maps network faults to business impact
type AlarmCorrelationEngine struct {
	rca *RCAEngine
}

func NewAlarmCorrelationEngine() *AlarmCorrelationEngine {
	return &AlarmCorrelationEngine{rca: &RCAEngine{}}
}

func (e *AlarmCorrelationEngine) Correlate(ctx context.Context, alarm *domain.Alarm) error {
	logger.Info("Correlating alarm",
		zap.String("id", alarm.ID),
		zap.String("resource", alarm.AlarmedResource.ID),
	)

	// In a real system, we would query topology and perform RCA
	// _, _ = e.rca.Analyze(ctx, []domain.Alarm{*alarm})

	if alarm.Severity == "Critical" {
		logger.Info("Critical outage detected. Triggering automated ticket generation.")
	}

	return nil
}

// RCAEngine implements heuristic-based Root Cause Analysis
type RCAEngine struct{}

func (e *RCAEngine) Analyze(ctx context.Context, alarms []domain.Alarm) (*domain.Alarm, error) {
	if len(alarms) == 0 {
		return nil, nil
	}

	logger.Info("Performing RCA on alarm group", zap.Int("count", len(alarms)))

	var rootCause *domain.Alarm
	for _, alarm := range alarms {
		if alarm.AlarmedResource.Type == "OLT" || alarm.AlarmedResource.Type == "CoreRouter" {
			rootCause = &alarm
			break
		}
	}

	if rootCause == nil {
		rootCause = &alarms[0]
	}

	logger.Info("RCA identified root cause", zap.String("id", rootCause.ID))
	return rootCause, nil
}
