package news

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/news"
)

type Endpoints struct {
	ListNewsByGenreEndpoint endpoint.Endpoint
	ListNewsEndpoint        endpoint.Endpoint
}

func MakeEndpoint(serv NewsManagement) Endpoints {
	return Endpoints{
		ListNewsEndpoint:        MakeGetNewsEndpoint(serv),
		ListNewsByGenreEndpoint: MakeGetNewsByGenreEndpoint(serv),
	}
}

func MakeGetNewsEndpoint(s NewsManagement) endpoint.Endpoint {
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
func MakeGetNewsByGenreEndpoint(s NewsManagement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ListNewsByGenreRequest)
		resp, err := s.ListNewsByGenre(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func DecodeGetNewsRequest(ctx context.Context, request interface{}) (interface{}, error) {
	fmt.Println("called this")
	req := request.(*pb.ListNewsRequest)
	return req, nil
}

func EncodeGetNewsResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.ListNewsResponse)
	return res, nil
}

//GetNewsByGenreRequest
func DecodeGetNewsByGenreRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ListNewsByGenreRequest)
	return req, nil
}

func EncodeGetNewsByGenreResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.ListNewsResponse)
	return res, nil
}
