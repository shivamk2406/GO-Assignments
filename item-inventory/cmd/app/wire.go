//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/shivamk2406/item-inventory/internal/service"
	"github.com/shivamk2406/item-inventory/internal/service/item"
)

var (
	ProviderSet wire.ProviderSet = wire.NewSet(
		service.ProviderConfig,
		service.ProviderDB,
		service.ProviderRepo,
	)
)

func Wire() *item.Repository {
	panic(wire.Build(ProviderSet))
}
