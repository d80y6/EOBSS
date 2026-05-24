package service

import (
	"context"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/kafka"
	"github.com/telcoflow/telcoflow/services/mediation-service/internal/domain"
	"time"
	"encoding/json"
)

type UsageTransformer struct {
	producer *kafka.Producer
}

func NewUsageTransformer(producer *kafka.Producer) *UsageTransformer {
	return &UsageTransformer{producer: producer}
}

func (t *UsageTransformer) Transform(ctx context.Context, rawData map[string]interface{}) (*domain.UsageEvent, error) {
	// TMF635 alignment
	event := &domain.UsageEvent{
		ID:        rawData["flow_id"].(string),
		UsageType: "Data",
		Quantity:  rawData["bytes"].(float64),
		Unit:      "Bytes",
		Timestamp: time.Now(),
		Source:    rawData["router_id"].(string),
	}

	// Map IP address to service_id via Inventory lookup or cache
	if serviceID, ok := rawData["service_id"].(string); ok {
		event.ServiceID = serviceID
	}

	// Forward to Kafka for Billing consumption
	msg, _ := json.Marshal(event)
	_ = t.producer.PublishEvent(ctx, event.ID, msg)

	return event, nil
}
