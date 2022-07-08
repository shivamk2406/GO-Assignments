package kproducer

import (
	"context"
	"log"

	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/kafka/producer"
)

type UserProducer interface {
	Produce(ctx context.Context, value *pb.ListActiveUsersResponse) error
}

type userProducer struct {
	Producer *producer.Producer
}

func NewUserProducer(p *producer.Producer) UserProducer {
	return userProducer{Producer: p}
}

func (up userProducer) Produce(ctx context.Context, value *pb.ListActiveUsersResponse) error {
	log.Println(value.ActiveUsers)
	return up.Producer.Produce(ctx, value)
}
