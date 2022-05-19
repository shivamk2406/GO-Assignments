//go:build wireinject
// +build wireinject

package app

import (
	"github.com/go-kit/log"
	"github.com/google/wire"
	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	"github.com/shivamk2406/newsletter-subscriptions/internal/kproducer"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/subscriptions"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/users"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/database"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/kafka/producer"
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

func initializeRegistry(db *gorm.DB, log log.Logger) *service.Registry {
	wire.Build(
		initializeActiveUserProducer,
		users.NewUsersRepo,
		users.UserManagementService,
		news.NewNewsRepo,
		news.NewsManagementService,
		subscriptions.NewSubscriptionRepo,
		subscriptions.NewSubscriptionService,
		service.ServiceRegistry)
	return &service.Registry{}
}
