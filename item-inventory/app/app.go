package app

import (
	"github.com/shivamk2406/item-inventory/internal/service"
)

func Start() error {
	err, db := service.NewRepo()
	if err != nil {
		return err
	}

	items, err := db.GetInventoryItem()
	if err != nil {
		return err
	}

	service.ProducerConsumerUtil(items)

	return nil
}
