package app

import (
	"github.com/shivamk2406/item-inventory/internal/service"
)

func Start() error {
	err := service.Init()
	if err != nil {
		return err
	}
	return nil
}
