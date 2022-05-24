//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/wire"
	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	"github.com/shivamk2406/newsletter-subscriptions/internal/kproducer"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/mail"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	subscriptions "github.com/shivamk2406/newsletter-subscriptions/internal/service/subscription"
	user "github.com/shivamk2406/newsletter-subscriptions/internal/service/user"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/database"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/kafka/producer"
	mailpkg "github.com/shivamk2406/newsletter-subscriptions/pkg/mail"
	"gorm.io/gorm"
)

func InitializeConfig() (config.Config, error) {
	wire.Build(config.LoadDatabaseConfig)
	return config.Config{}, nil
}

func initializeDB(conf config.Config) (*gorm.DB, func(), error) {
	wire.Build(database.Open)
	return &gorm.DB{}, func() {}, nil
}

func initializeActiveUserProducer() kproducer.UserProducer {
	panic(wire.Build(
		config.LoadProducerConfig,
		producer.NewProducer,
		kproducer.NewUserProducer,
	))
}

func initializeRegistry(ctx context.Context, db *gorm.DB, log log.Logger) *service.Registry {
	wire.Build(
		initializeActiveUserProducer,
		user.NewUserRepo,
		user.UserManagementService,
		news.NewNewsRepo,
		news.NewsManagementService,
		subscriptions.NewSubscriptionRepo,
		subscriptions.NewSubscriptionService,
		service.ServiceRegistry,
		config.LoadMailService,
		mailpkg.NewMailConn,
		mail.NewMailService,
	)
	return &service.Registry{}
}
