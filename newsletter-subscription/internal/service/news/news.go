package news

import (
	"context"

	"github.com/go-kit/log"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/news"
)

type NewsManagement interface {
	ListNews(ctx context.Context, in *pb.ListNewsRequest) (*pb.ListNewsResponse, error)
	ListNewsByGenre(ctx context.Context, in *pb.ListNewsByGenreRequest) (*pb.ListNewsResponse, error)
}

type NewsManagementServer struct {
	log  log.Logger
	repo NewsDB
}

func NewsManagementService(repo NewsDB, logger log.Logger) NewsManagement {
	return &NewsManagementServer{repo: repo, log: logger}
}

func (r NewsManagementServer) ListNews(ctx context.Context, in *pb.ListNewsRequest) (*pb.ListNewsResponse, error) {
	model := GetNewsRequest{Subsid: in.Subsid}
	response, err := r.repo.getNews(ctx, model)
	if err != nil {
		return &pb.ListNewsResponse{}, err
	}
	var newsString []*pb.News
	for _, val := range response.Newss {
		newsString = append(newsString, &pb.News{Heading: val.heading, Description: val.description})
	}
	return &pb.ListNewsResponse{News: newsString}, nil
}

func (r NewsManagementServer) ListNewsByGenre(ctx context.Context, in *pb.ListNewsByGenreRequest) (*pb.ListNewsResponse, error) {
	model := GetNewsByGenreRequest{Genre: in.Genre}
	response, err := r.repo.getNewsByGenre(ctx, model)
	if err != nil {
		return &pb.ListNewsResponse{}, err
	}
	var newsString []*pb.News
	for _, val := range response.Newss {
		newsString = append(newsString, &pb.News{Heading: val.heading, Description: val.description})
	}
	return &pb.ListNewsResponse{News: newsString}, nil
}
