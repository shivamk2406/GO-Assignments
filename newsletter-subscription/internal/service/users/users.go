package users

import (
	"context"

	"github.com/go-kit/log"

	"github.com/shivamk2406/newsletter-subscriptions/internal/kproducer"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
)

type UserManagement interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error)
	AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error)
	ListActiveUsers(ctx context.Context, in *pb.ListActiveUsersRequest) (*pb.ListActiveUsersResponse, error)
}

type UserManagementServer struct {
	log                  log.Logger
	repo                 UsersDB
	activeUsersProducers kproducer.UserProducer
}

func UserManagementService(repo UsersDB, logger log.Logger, uproducer kproducer.UserProducer) UserManagement {
	return &UserManagementServer{
		repo:                 repo,
		log:                  logger,
		activeUsersProducers: uproducer,
	}
}

func (r UserManagementServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	model := CreateUserRequest{Name: in.Name, Email: in.Email}
	user, err := r.repo.createUser(ctx, model)
	if err != nil {
		return &pb.User{}, err
	}
	return &pb.User{Name: user.Email, Email: user.Email, Active: false}, nil
}

func (r UserManagementServer) AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	//w := log.NewSyncWriter(os.Stderr)
	//logger := log.NewLogfmtLogger(w)
	model := AuthenticateUserRequest{Email: in.Email}
	//logger.Log("Inside User Request Generated for db %v", model)

	response, err := r.repo.authenticateUser(ctx, model)
	if err != nil {
		return &pb.AuthenticateUserResponse{}, err
	}

	user := pb.User{Name: response.User.Name, Email: response.User.Email, Active: response.User.Active}
	return &pb.AuthenticateUserResponse{IsAuthenticated: response.IsAuthenticated, User: &user}, nil
}

func (r UserManagementServer) ListActiveUsers(ctx context.Context, in *pb.ListActiveUsersRequest) (*pb.ListActiveUsersResponse, error) {
	model := ListActiveUsers{}
	response, err := r.repo.listActiveUsers(ctx, model)
	if err != nil {
		return &pb.ListActiveUsersResponse{}, err
	}

	var activeUsers []*pb.User
	for _, val := range response.ActiveUsers {
		activeUsers = append(activeUsers, &pb.User{Name: val.Name, Email: val.Email, Active: val.Active})
	}

	err = r.activeUsersProducers.Produce(ctx, &pb.ListActiveUsersResponse{ActiveUsers: activeUsers})
	if err != nil {
		return &pb.ListActiveUsersResponse{}, err
	}
	return &pb.ListActiveUsersResponse{ActiveUsers: activeUsers}, nil

}
