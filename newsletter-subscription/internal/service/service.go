package service

import (
	"context"
	"fmt"
	"log"

	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	newsletter "github.com/shivamk2406/newsletter-subscriptions/internal/proto/news"
	subspb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/subscriptions"
	userpb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
	mailsvc "github.com/shivamk2406/newsletter-subscriptions/internal/service/mail"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	subscription "github.com/shivamk2406/newsletter-subscriptions/internal/service/subscription"
	user "github.com/shivamk2406/newsletter-subscriptions/internal/service/user"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/kafka/consumer"
	"google.golang.org/protobuf/encoding/protojson"
)

type Registry struct {
	SubscriptionService subscription.SubscriptionManagement
	NewsService         news.NewsManagement
	UsersService        user.UserManagement
	MailService         mailsvc.MailService
}

var unmarshaller = &protojson.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: true,
}

func ServiceRegistry(SubsServ subscription.SubscriptionManagement, NewsServ news.NewsManagement,
	UsersServ user.UserManagement, MailServ mailsvc.MailService) *Registry {
	return &Registry{
		SubscriptionService: SubsServ,
		NewsService:         NewsServ,
		UsersService:        UsersServ,
		MailService:         MailServ,
	}
}

func (r Registry) CronService(ctx context.Context, consumer consumer.Consumer) {
	_, err := r.UsersService.ListActiveUsers(ctx, &userpb.ListActiveUsersRequest{})
	if err != nil {
		log.Println(err)
	}

	go consumer.Start(ctx, r.postConsume)
}

func (r Registry) postConsume(ctx context.Context, b []byte) error {
	cfg := config.LoadMailService()
	fmt.Println("postConsume")
	req := userpb.ListActiveUsersResponse{}
	if err := unmarshaller.Unmarshal(b, &req); err != nil {
		log.Println(err)
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
		mail := r.MailService.CreateMail(val.Email, cfg.Username, "Your Daily Feed", newsCollection)
		r.MailService.SendMail(mail)
	}
	return nil
}
