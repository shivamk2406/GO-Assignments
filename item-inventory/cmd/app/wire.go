//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"

	"github.com/shivamk2406/item-inventory/internal/config"
	"github.com/shivamk2406/item-inventory/internal/service/item"
	"github.com/shivamk2406/item-inventory/pkg/database"
)

func InitializeEvent() (*item.Repository, error) {
	panic(wire.Build(wire.NewSet(
		config.InitializeConfig,
		database.InitializeDB,
		item.InitializeRepo,
	)))
	return &item.Repository{}, nil
}
