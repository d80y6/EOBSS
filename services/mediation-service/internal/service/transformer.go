package service

import (
	"context"
	"github.com/telcoflow/telcoflow/services/mediation-service/internal/domain"
	"time"
)

type UsageTransformer struct {
	// kafkaProducer KafkaProducer
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
	// event.ServiceID = t.lookupServiceID(rawData["src_ip"].(string))

	return event, nil
}
