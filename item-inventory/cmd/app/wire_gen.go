// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/shivamk2406/item-inventory/internal/config"
	"github.com/shivamk2406/item-inventory/internal/service/item"
	"github.com/shivamk2406/item-inventory/pkg/database"
)

// Injectors from wire.go:

func InitializeEvent() *item.Repository {
	configConfig := config.InitializeConfig()
	db := database.InitializeDB(configConfig)
	repository := item.InitializeRepo(db)
	return repository
}
