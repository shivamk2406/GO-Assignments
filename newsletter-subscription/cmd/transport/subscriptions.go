package transport

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/subscriptions"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service"
	subscriptions "github.com/shivamk2406/newsletter-subscriptions/internal/service/subscription"
)

type subscriptionsServer struct {
	pb.UnimplementedSubscriptionManagementServiceServer
	ListPlansHandler          grpc.Handler
	GetSubscriptionHandler    grpc.Handler
	CreateSubscriptionHandler grpc.Handler
}

func NewSubscriptionServer(ctx context.Context, reg service.Registry) pb.SubscriptionManagementServiceServer {
	return &subscriptionsServer{
		ListPlansHandler: grpc.NewServer(subscriptions.MakeGetPlansEndpoint(reg.SubscriptionService),
			subscriptions.DecodeGetPlansRequest,
			subscriptions.EncodeGetPlansResponse),
		GetSubscriptionHandler: grpc.NewServer(subscriptions.MakeGetSubscriptionEndpoint(reg.SubscriptionService),
			subscriptions.DecodeGetSubscriptionRequest,
			subscriptions.EncodeGetSubscriptionResponse),
		CreateSubscriptionHandler: grpc.NewServer(subscriptions.MakeSetSubscriptionEndpoint(reg.SubscriptionService),
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
