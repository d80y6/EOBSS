package mediation

import (
	"context"
	"encoding/json"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/rating"
	"go.uber.org/zap"
)

// UsageMediator orchestrates the collection, rating, and persistence of usage events
type UsageMediator struct {
	engine *rating.RatingEngine
	// kafkaConsumer *kafka.Reader
	// clickhouseDB  *sql.DB
}

func NewUsageMediator(engine *rating.RatingEngine) *UsageMediator {
	return &UsageMediator{
		engine: engine,
	}
}

func (m *UsageMediator) ProcessEvent(ctx context.Context, rawEvent []byte) error {
	var record domain.UsageRecord
	if err := json.Unmarshal(rawEvent, &record); err != nil {
		logger.Error("Failed to unmarshal usage event", zap.Error(err))
		return err
	}

	// 1. Enrich & Validate
	// (Check if service is active in Inventory)

	// 2. Rate the usage
	if err := m.engine.RateUsage(ctx, &record); err != nil {
		logger.Error("Failed to rate usage", zap.String("id", record.ID), zap.Error(err))
		return err
	}

	// 3. Persist to ClickHouse (Long-term storage for analytics and invoicing)
	if err := m.persistToClickHouse(ctx, &record); err != nil {
		return err
	}

	// 4. Update real-time balance in Redis
	// (Logic for prepaid wallets)

	logger.Info("Usage record processed successfully",
		zap.String("id", record.ID),
		zap.Float64("amount", record.RatedAmount),
	)

	return nil
}

func (m *UsageMediator) persistToClickHouse(ctx context.Context, record *domain.UsageRecord) error {
	// Implementation for ClickHouse batch insertion
	// INSERT INTO cdr (id, usage_type, quantity, amount, timestamp) VALUES (...)
	return nil
}
