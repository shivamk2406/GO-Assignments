package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
)

func NewKafkaConn(cfg *config.Config) (*kafka.Conn, error) {
	return kafka.DialContext(context.Background(), "tcp", cfg.Kafka.Brokers)
}
