package users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
)

type Endpoints struct {
	CreateUserEndpoint       endpoint.Endpoint
	AuthenticateUserEndpoint endpoint.Endpoint
	ListActiveUsersEndpoint  endpoint.Endpoint
}

func MakeEndpoint(serv UserManagement) Endpoints {
	return Endpoints{
		CreateUserEndpoint:       MakeCreateUserEndpoint(serv),
		AuthenticateUserEndpoint: MakeAuthenticateUserEndpoint(serv),
		ListActiveUsersEndpoint:  MakeListActiveUserEndpoinr(serv),
	}
}

func MakeCreateUserEndpoint(s UserManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateUserRequest)
		resp, err := s.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func MakeAuthenticateUserEndpoint(s UserManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.AuthenticateUserRequest)
		resp, err := s.AuthenticateUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func MakeListActiveUserEndpoinr(s UserManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ListActiveUsersRequest)
		resp, err := s.ListActiveUsers(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func DecodeCreateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return req, nil
}

func EncodeCreateUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.User)
	return resp, nil
}

func DecodeAuthenticateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AuthenticateUserRequest)
	return req, nil
}

func EncodeAuthenticateUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.AuthenticateUserResponse)
	user := pb.User{Name: res.User.Name, Email: res.User.Email, Active: res.User.Active}
	return &pb.AuthenticateUserResponse{IsAuthenticated: res.IsAuthenticated, User: &user}, nil
}

func DecodeListActiveUsersRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ListActiveUsersRequest)
	return req, nil
}

func EncodeListActiveUsersResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.ListActiveUsersResponse)
	return res, nil
}
