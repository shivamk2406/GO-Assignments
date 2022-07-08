package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type ConsumerConfig struct {
	BootstrapServers []string `json:"bootstrap_servers,omitempty"`
	Topic            string   `json:"topic,omitempty"`
	Group            string   `json:"group,omitempty"`
}

type Consumer struct {
	reader *kafka.Reader
	config *kafka.ReaderConfig
}

type ConsumeFunc func(c context.Context, b []byte) error

func NewConsumer(_ context.Context, cfg ConsumerConfig) (*Consumer, error) {
	consumer := &Consumer{
		config: &kafka.ReaderConfig{
			Brokers: cfg.BootstrapServers,
			GroupID: cfg.Group,
			Topic:   cfg.Topic,
		},
	}
	return consumer, nil
}

func (c Consumer) Start(ctx context.Context, postconsume ConsumeFunc) {
	fmt.Println("Consumer started")

	defer func() {
		if e := recover(); e != nil {
			log.Printf("recovered from panic due to %v", e)
		}
	}()
	c.reader = kafka.NewReader(*c.config)
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("errRead")
			panic(err)
		}
		err = postconsume(ctx, msg.Value)
		if err != nil {
			log.Println(err)
		}
	}

}
