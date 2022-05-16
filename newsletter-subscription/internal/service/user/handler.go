package user

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
)

type Endpoints struct {
	CreateUserEndpoint        endpoint.Endpoint
	ListPlansEndpoint         endpoint.Endpoint
	AuthenticateUserEndpoint  endpoint.Endpoint
	GetSubscriptionEnpoint    endpoint.Endpoint
	CreateSubscrptionEndpoint endpoint.Endpoint
	ListNewsByGenreEndpoint   endpoint.Endpoint
	ListNewsEndpoint          endpoint.Endpoint
}

func MakeEndpoint(serv UserManagement) Endpoints {
	return Endpoints{CreateUserEndpoint: MakeCreateUserEndpoint(serv),
		ListPlansEndpoint:         MakeGetPlansEndpoint(serv),
		AuthenticateUserEndpoint:  MakeAuthenticateUserEndpoint(serv),
		GetSubscriptionEnpoint:    MakeGetSubscriptionEndpoint(serv),
		CreateSubscrptionEndpoint: MakeSetSubscriptionEndpoint(serv),
		ListNewsEndpoint:          MakeGetNewsEndpoint(serv),
		ListNewsByGenreEndpoint:   MakeGetNewsByGenreEndpoint(serv),
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

//Make endpoint for Getting All Plans
func MakeGetPlansEndpoint(s UserManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ListPlansRequest)
		resp, err := s.ListPlans(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

//Make endpoint for Authenticating User
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

//Make endpoint for Getting Subscriptions
func MakeGetSubscriptionEndpoint(s UserManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SubscriptionRequest)
		resp, err := s.GetSubscription(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

//Make endpoint for Setting Subscriptions
func MakeSetSubscriptionEndpoint(s UserManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateSubscriptionRequest)
		resp, err := s.CreateSubscription(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

//Make endpoint for Getting News
func MakeGetNewsEndpoint(s UserManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ListNewsRequest)
		resp, err := s.ListNews(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

//Make endpoint for Getting News by Genre
func MakeGetNewsByGenreEndpoint(s UserManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ListNewsByGenreRequest)
		resp, err := s.ListNewsByGenre(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

//Create user
func DecodeCreateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return req, nil
}

func EncodeCreateUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.User)
	return resp, nil
}

//Get Plans
func DecodeGetPlansRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ListPlansRequest)
	return req, nil
}

func EncodeGetPlansResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Plans)
	return resp, nil
}

//Authenticate User
func DecodeAuthenticateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AuthenticateUserRequest)
	return req, nil
}

func EncodeAuthenticateUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.AuthenticateUserResponse)
	user := pb.User{Name: res.User.Name, Email: res.User.Email, Active: res.User.Active}
	return &pb.AuthenticateUserResponse{IsAuthenticated: res.IsAuthenticated, User: &user}, nil
}

//GetSubscription
func DecodeGetSubscriptionRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SubscriptionRequest)
	return req, nil
}

func EncodeGetSubscriptionResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.Plan)
	// genres := response.([]*pb.Genre)
	// for _, val := range res.Genres {
	// 	gen := pb.Genre{Name: val.Name}
	// 	genres = append(genres, &gen)
	// }
	return res, nil
}

//SetSubscription
func DecodeSetSubscriptionRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateSubscriptionRequest)
	return req, nil
}

func EncodeSetSubscriptionResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.SubscriptionResponse)
	return res, nil
}

//GetNews
func DecodeGetNewsRequest(ctx context.Context, request interface{}) (interface{}, error) {
	fmt.Println("called this")
	req := request.(*pb.ListNewsRequest)
	return req, nil
}

func EncodeGetNewsResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.News)
	return res, nil
}

//GetNewsByGenreRequest
func DecodeGetNewsByGenreRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ListNewsByGenreRequest)
	return req, nil
}

func EncodeGetNewsByGenreResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.News)
	return res, nil
}
