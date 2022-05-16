//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	"github.com/shivamk2406/newsletter-subscriptions/internal/pkg/database"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/user"
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

func initializeRepo(db *gorm.DB) *user.Repository {
	wire.Build(user.NewRepo)
	return &user.Repository{}
}
