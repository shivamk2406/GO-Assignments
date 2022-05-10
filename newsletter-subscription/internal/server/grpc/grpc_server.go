package server

import (
	"context"
	"log"
	"net"

	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	port = ":50051"
)

type UserServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserReponse, error) {
	log.Printf("Received: %v", in.Subsid)
	log.Printf("Created User: %s %s %d %d %d %v", in.Name, in.Email, 1, 30, in.Subsid, timestamppb.Now().AsTime())
	return &pb.CreateUserReponse{Name: in.Name, Email: in.Email, Active: 1, Validity: 30, Subsid: in.Subsid, Starttime: timestamppb.Now()}, nil
}

func (s *UserServer) AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserRespone, error) {
	log.Printf("Received : %v", in.Email)
	return &pb.AuthenticateUserRespone{IsAuthenticated: true}, nil
}

func RunServer(ctx context.Context, repo pb.UserManagementServer) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, repo)
	log.Printf("server listening at %v", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
		return err
	}
	return nil
}
