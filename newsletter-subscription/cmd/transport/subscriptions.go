package transport

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/subscriptions"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/subscriptions"
)

type subscriptionsServer struct {
	pb.UnimplementedSubscriptionManagementServiceServer
	ListPlansHandler          grpc.Handler
	GetSubscriptionHandler    grpc.Handler
	CreateSubscriptionHandler grpc.Handler
}

func NewSubscriptionServer(e subscriptions.Endpoints) pb.SubscriptionManagementServiceServer {
	return &subscriptionsServer{
		ListPlansHandler: grpc.NewServer(e.ListPlansEndpoint,
			subscriptions.DecodeGetPlansRequest,
			subscriptions.EncodeGetPlansResponse),
		GetSubscriptionHandler: grpc.NewServer(e.GetSubscriptionEnpoint,
			subscriptions.DecodeGetSubscriptionRequest,
			subscriptions.EncodeGetSubscriptionResponse),
		CreateSubscriptionHandler: grpc.NewServer(e.CreateSubscrptionEndpoint,
			subscriptions.DecodeSetSubscriptionRequest,
			subscriptions.EncodeSetSubscriptionResponse),
	}
}

func (s *subscriptionsServer) ListPlans(ctx context.Context, in *pb.ListPlansRequest) (*pb.Plans, error) {
	_, res, err := s.ListPlansHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Plans), nil

}

func (s *subscriptionsServer) GetSubscription(ctx context.Context, in *pb.SubscriptionRequest) (*pb.Plan, error) {
	_, res, err := s.GetSubscriptionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Plan), nil
}

func (s *subscriptionsServer) CreateSubscription(ctx context.Context, in *pb.CreateSubscriptionRequest) (*pb.SubscriptionResponse, error) {
	_, res, err := s.CreateSubscriptionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SubscriptionResponse), nil
}
