package service

import (
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/subscriptions"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/users"
)

type Registry struct {
	SubscriptionService subscriptions.SubscriptionManagement
	NewsService         news.NewsManagement
	UsersService        users.UserManagement
}

func ServiceRegistry(SubsServ subscriptions.SubscriptionManagement,
	NewsServ news.NewsManagement,
	UsersServ users.UserManagement) *Registry {
	return &Registry{SubscriptionService: SubsServ,
		NewsService:  NewsServ,
		UsersService: UsersServ}
}
