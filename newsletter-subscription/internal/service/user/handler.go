package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
)

//Make endpoint for create user
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
