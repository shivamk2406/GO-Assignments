package subscriptions

import (
	"context"

	"github.com/go-kit/log"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/subscriptions"
)

type SubscriptionManagement interface {
	ListPlans(ctx context.Context, in *pb.ListPlansRequest) (*pb.Plans, error)
	GetSubscription(ctx context.Context, in *pb.SubscriptionRequest) (*pb.Plan, error)
	CreateSubscription(ctx context.Context, in *pb.CreateSubscriptionRequest) (*pb.SubscriptionResponse, error)
}

type SubscriptionManagementServer struct {
	log  log.Logger
	repo SubscriptionDB
}

func NewSubscriptionService(repos SubscriptionDB, logger log.Logger) SubscriptionManagement {
	return &SubscriptionManagementServer{repo: repos, log: logger}
}

func (r SubscriptionManagementServer) ListPlans(ctx context.Context, in *pb.ListPlansRequest) (*pb.Plans, error) {
	model := GetPlansRequests{}
	plans, err := r.repo.getPlans(ctx, model)
	if err != nil {
		return &pb.Plans{}, err
	}

	var subs []*pb.Plan
	for _, val := range plans.Subs {
		sub := pb.Plan{Id: val.Id, Name: val.Name, Validity: val.Validity}
		var genres []*pb.Genre
		for _, val1 := range val.Genres {
			genre := pb.Genre{Name: val1.Name}
			genres = append(genres, &genre)
		}
		sub.Genres = genres
		subs = append(subs, &sub)
	}
	return &pb.Plans{Subs: subs}, nil
}

func (r SubscriptionManagementServer) GetSubscription(ctx context.Context, in *pb.SubscriptionRequest) (*pb.Plan, error) {
	model := SubscriptionRequest{Email: in.Email}
	response, err := r.repo.getSubscription(ctx, model)
	if err != nil {
		return &pb.Plan{}, err
	}
	var genres []*pb.Genre
	for _, val := range response.Genres {
		genre := pb.Genre{Name: val.Name}
		genres = append(genres, &genre)
	}
	return &pb.Plan{Id: response.Id, Name: response.Name, Validity: response.Validity, Genres: genres}, nil
}

func (r SubscriptionManagementServer) CreateSubscription(ctx context.Context, in *pb.CreateSubscriptionRequest) (*pb.SubscriptionResponse, error) {
	model := SetSubscriptionRequest{Email: in.Email, Subsid: in.Subsid}
	response, err := r.repo.setSubscription(ctx, model)
	if err != nil {
		return &pb.SubscriptionResponse{}, err
	}
	return &pb.SubscriptionResponse{Email: in.Email, Active: true, Starttime: &response.Starttime, Validity: response.Validity}, nil
}
