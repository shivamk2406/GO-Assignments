package transport

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/news"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
)

type newsServer struct {
	pb.UnimplementedNewsServiceServer
	ListNewsByGenreHandler grpc.Handler
	ListNewsHandler        grpc.Handler
}

func NewNewsServer(e news.Endpoints) pb.NewsServiceServer {
	return &newsServer{
		ListNewsHandler: grpc.NewServer(e.ListNewsEndpoint,
			news.DecodeGetNewsRequest,
			news.EncodeGetNewsResponse),
		ListNewsByGenreHandler: grpc.NewServer(e.ListNewsByGenreEndpoint,
			news.DecodeGetNewsByGenreRequest,
			news.EncodeGetNewsByGenreResponse),
	}
}

func (s *newsServer) ListNews(ctx context.Context, in *pb.ListNewsRequest) (*pb.ListNewsResponse, error) {
	_, res, err := s.ListNewsHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListNewsResponse), nil
}

func (s *newsServer) ListNewsByGenre(ctx context.Context, in *pb.ListNewsByGenreRequest) (*pb.ListNewsResponse, error) {
	_, res, err := s.ListNewsByGenreHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListNewsResponse), nil
}
