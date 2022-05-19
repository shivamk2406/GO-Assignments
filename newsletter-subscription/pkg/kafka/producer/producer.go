package producer

import (
	"context"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type ProducerConfig struct {
	BootstrapServers []string `json:"bootstrap_servers,omitempty"`
	Topic            string   `json:"topic,omitempty"`
}

type Producer struct {
	marshaller      protojson.MarshalOptions
	protoMarshaller proto.MarshalOptions
	protojson       bool
	writer          *kafka.Writer
}

type ProducerOption func(*Producer)

func WithProtoMarshaller(m proto.MarshalOptions) ProducerOption {
	return func(p *Producer) {
		p.protoMarshaller = m
	}
}

func WithProto() ProducerOption {
	return func(p *Producer) {
		p.protojson = false
	}
}

func WithProtoJSONMarshalOpts(m protojson.MarshalOptions) ProducerOption {
	return func(p *Producer) {
		p.marshaller = m
	}
}

func NewProducer(cfg ProducerConfig) *Producer {

	producer := &Producer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(cfg.BootstrapServers...),
			Topic:        cfg.Topic,
			Async:        false,
			RequiredAcks: kafka.RequireAll,
		},
		marshaller:      protojson.MarshalOptions{},
		protoMarshaller: proto.MarshalOptions{},
		protojson:       true,
	}

	return producer
}

func (p *Producer) Produce(ctx context.Context, pb proto.Message) error {
	var val []byte
	var err error

	if p.protojson {
		val, err = p.marshaller.Marshal(pb)
	} else {
		val, err = p.protoMarshaller.Marshal(pb)
	}
	if err != nil {
		return errors.WithMessage(err, "producer: failed to marshal message")
	}

	// GrpcMetadataKey to KafkaHeaders Key
	msg := kafka.Message{
		Value: val,
		// TODO: add trace id
	}

	err = p.writer.WriteMessages(ctx, msg)
	if err != nil {
		return errors.WithMessagef(err, "producer: failed publish msg to topic %s", p.writer.Topic)
	}
	return nil
}
