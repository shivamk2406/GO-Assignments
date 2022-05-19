package transport

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/users"
)

type userServer struct {
	pb.UnimplementedUserManagementServiceServer
	CreateUserHandler       grpc.Handler
	AuthenticateUserHandler grpc.Handler
	ListActiveUsersHandler  grpc.Handler
}

func NewUserGrpcServer(e users.Endpoints) pb.UserManagementServiceServer {
	return &userServer{
		CreateUserHandler: grpc.NewServer(e.CreateUserEndpoint,
			users.DecodeCreateUserRequest,
			users.EncodeCreateUserResponse,
		),
		AuthenticateUserHandler: grpc.NewServer(e.AuthenticateUserEndpoint,
			users.DecodeAuthenticateUserRequest,
			users.EncodeAuthenticateUserResponse),
		ListActiveUsersHandler: grpc.NewServer(e.ListActiveUsersEndpoint,
			users.DecodeListActiveUsersRequest,
			users.EncodeListActiveUsersResponse),
	}
}

func (s *userServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	_, res, err := s.CreateUserHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.User), nil
}

func (s *userServer) AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	_, res, err := s.AuthenticateUserHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.AuthenticateUserResponse), nil
}

func (s *userServer) ListActiveUsers(ctx context.Context, in *pb.ListActiveUsersRequest) (*pb.ListActiveUsersResponse, error) {
	_, res, err := s.ListActiveUsersHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListActiveUsersResponse), nil
}
