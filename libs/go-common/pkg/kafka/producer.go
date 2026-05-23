package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"go.uber.org/zap"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Producer) PublishEvent(ctx context.Context, key string, value []byte) error {
	logger.Info("Publishing event to Kafka", zap.String("topic", p.writer.Topic), zap.String("key", key))
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: value,
	})
}
