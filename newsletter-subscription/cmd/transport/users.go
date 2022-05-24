package transport

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service"
	user "github.com/shivamk2406/newsletter-subscriptions/internal/service/user"
)

type userServer struct {
	pb.UnimplementedUserManagementServiceServer
	CreateUserHandler       grpc.Handler
	AuthenticateUserHandler grpc.Handler
	ListActiveUsersHandler  grpc.Handler
}

func NewUserGrpcServer(ctx context.Context, reg service.Registry) pb.UserManagementServiceServer {
	return &userServer{
		CreateUserHandler: grpc.NewServer(user.MakeCreateUserEndpoint(reg.UsersService),
			user.DecodeCreateUserRequest,
			user.EncodeCreateUserResponse,
		),
		AuthenticateUserHandler: grpc.NewServer(user.MakeAuthenticateUserEndpoint(reg.UsersService),
			user.DecodeAuthenticateUserRequest,
			user.EncodeAuthenticateUserResponse),
		ListActiveUsersHandler: grpc.NewServer(user.MakeListActiveUserEndpoint(reg.UsersService),
			user.DecodeListActiveUsersRequest,
			user.EncodeListActiveUsersResponse),
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
