package service

import (
	"context"
	"fmt"

	newsletter "github.com/shivamk2406/newsletter-subscriptions/internal/proto/news"
	subspb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/subscriptions"
	userpb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/mail"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/subscriptions"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/users"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/kafka/consumer"
	"google.golang.org/protobuf/encoding/protojson"
)

type Registry struct {
	SubscriptionService subscriptions.SubscriptionManagement
	NewsService         news.NewsManagement
	UsersService        users.UserManagement
}

var unmarshaller = &protojson.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: true,
}

func ServiceRegistry(SubsServ subscriptions.SubscriptionManagement,
	NewsServ news.NewsManagement,
	UsersServ users.UserManagement) *Registry {
	return &Registry{SubscriptionService: SubsServ,
		NewsService:  NewsServ,
		UsersService: UsersServ}
}

func (r Registry) CronService(ctx context.Context, consumer consumer.Consumer) {
	_, err := r.UsersService.ListActiveUsers(ctx, &userpb.ListActiveUsersRequest{})
	if err != nil {
		fmt.Println(err)
	}

	go consumer.Start(ctx, r.postConsume)
}

func (r Registry) postConsume(ctx context.Context, b []byte) error {
	fmt.Println("postConsume")
	req := userpb.ListActiveUsersResponse{}
	if err := unmarshaller.Unmarshal(b, &req); err != nil {
		fmt.Println(err)
		return err
	}
	activeUsers := req.GetActiveUsers()
	var emails []string

	for _, val := range activeUsers {
		emails = append(emails, val.Email)
		subs, _ := r.SubscriptionService.GetSubscription(ctx, &subspb.SubscriptionRequest{Email: val.Email})
		newsCollectionResponse, _ := r.NewsService.ListNews(ctx, &newsletter.ListNewsRequest{Subsid: subs.Id})
		var newsCollection []news.SingleNews
		for _, val := range newsCollectionResponse.News {
			newsCollection = append(newsCollection, news.SingleNews{Heading: val.Heading, Description: val.Description})
		}
		mailService := mail.NewMailService(val.Email, newsCollection)
		mailService.SendMail()
	}
	return nil
}
