package transport

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/user"
)

type gRPCServer struct {
	pb.UnimplementedUserManagementServer
	CreateUserHandler         grpc.Handler
	ListPlansHandler          grpc.Handler
	AuthenticateUserHandler   grpc.Handler
	GetSubscriptionHandler    grpc.Handler
	CreateSubscriptionHandler grpc.Handler
	ListNewsByGenreHandler    grpc.Handler
	ListNewsHandler           grpc.Handler
}

func NewServer(e user.Endpoints) pb.UserManagementServer {
	return &gRPCServer{
		CreateUserHandler: grpc.NewServer(e.CreateUserEndpoint,
			user.DecodeCreateUserRequest,
			user.EncodeCreateUserResponse),
		ListPlansHandler: grpc.NewServer(e.ListPlansEndpoint,
			user.DecodeGetPlansRequest,
			user.EncodeGetPlansResponse),
		AuthenticateUserHandler: grpc.NewServer(e.AuthenticateUserEndpoint,
			user.DecodeAuthenticateUserRequest,
			user.EncodeAuthenticateUserResponse),
		GetSubscriptionHandler: grpc.NewServer(e.GetSubscriptionEnpoint,
			user.DecodeGetSubscriptionRequest,
			user.EncodeGetSubscriptionResponse),
		CreateSubscriptionHandler: grpc.NewServer(e.CreateSubscrptionEndpoint,
			user.DecodeSetSubscriptionRequest,
			user.EncodeSetSubscriptionResponse),
		ListNewsHandler: grpc.NewServer(e.ListNewsEndpoint,
			user.DecodeGetNewsRequest,
			user.EncodeGetNewsResponse),
		ListNewsByGenreHandler: grpc.NewServer(e.ListNewsByGenreEndpoint,
			user.DecodeGetNewsByGenreRequest,
			user.EncodeGetNewsByGenreResponse),
	}
}

func (s *gRPCServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	_, res, err := s.CreateUserHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.User), nil
}

func (s *gRPCServer) ListPlans(ctx context.Context, in *pb.ListPlansRequest) (*pb.Plans, error) {
	_, res, err := s.ListPlansHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Plans), nil

}
func (s *gRPCServer) AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	_, res, err := s.AuthenticateUserHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.AuthenticateUserResponse), nil
}
func (s *gRPCServer) GetSubscription(ctx context.Context, in *pb.SubscriptionRequest) (*pb.Plan, error) {
	_, res, err := s.GetSubscriptionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Plan), nil
}

func (s *gRPCServer) CreateSubscription(ctx context.Context, in *pb.CreateSubscriptionRequest) (*pb.SubscriptionResponse, error) {
	_, res, err := s.CreateSubscriptionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SubscriptionResponse), nil
}
func (s *gRPCServer) ListNews(ctx context.Context, in *pb.ListNewsRequest) (*pb.News, error) {
	_, res, err := s.ListNewsHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.News), nil
}

func (s *gRPCServer) ListNewsByGenre(ctx context.Context, in *pb.ListNewsByGenreRequest) (*pb.News, error) {
	_, res, err := s.ListNewsByGenreHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.News), nil
}
