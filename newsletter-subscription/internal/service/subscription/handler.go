package subscriptions

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/subscriptions"
)

type Endpoints struct {
	ListPlansEndpoint         endpoint.Endpoint
	GetSubscriptionEnpoint    endpoint.Endpoint
	CreateSubscrptionEndpoint endpoint.Endpoint
}

func MakeEndpoint(serv SubscriptionManagement) Endpoints {
	return Endpoints{
		ListPlansEndpoint:         MakeGetPlansEndpoint(serv),
		GetSubscriptionEnpoint:    MakeGetSubscriptionEndpoint(serv),
		CreateSubscrptionEndpoint: MakeSetSubscriptionEndpoint(serv),
	}
}

//Make endpoint for Getting All Plans
func MakeGetPlansEndpoint(s SubscriptionManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ListPlansRequest)
		resp, err := s.ListPlans(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func MakeGetSubscriptionEndpoint(s SubscriptionManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SubscriptionRequest)
		resp, err := s.GetSubscription(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func MakeSetSubscriptionEndpoint(s SubscriptionManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateSubscriptionRequest)
		resp, err := s.CreateSubscription(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func DecodeGetPlansRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ListPlansRequest)
	return req, nil
}

func EncodeGetPlansResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Plans)
	return resp, nil
}

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

func DecodeSetSubscriptionRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateSubscriptionRequest)
	return req, nil
}

func EncodeSetSubscriptionResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.SubscriptionResponse)
	return res, nil
}
