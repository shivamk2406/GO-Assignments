// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	"github.com/shivamk2406/newsletter-subscriptions/internal/pkg/database"
	"github.com/shivamk2406/newsletter-subscriptions/internal/proto"
	"github.com/shivamk2406/newsletter-subscriptions/internal/user"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeConfig() (config.Config, error) {
	configConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		return config.Config{}, err
	}
	return configConfig, nil
}

func initializeDB(conf config.Config) (*gorm.DB, func(), error) {
	db, cleanup, err := database.Open(conf)
	if err != nil {
		return nil, nil, err
	}
	return db, func() {
		cleanup()
	}, nil
}

func initializeRepo(db *gorm.DB) newsletter.UserManagementServer {
	userManagementServer := user.NewRepo(db)
	return userManagementServer
}