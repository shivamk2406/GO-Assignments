package service

import "github.com/shivamk2406/item-inventory/database"

func Init() error {
	config := LoadAppConfig()

	db, err := database.Open(config)
	if err != nil {
		return err
	}

	err, repo := NewRepo(db)
	if err != nil {
		return err
	}

	items, err := repo.GetInventoryItem()
	if err != nil {
		return err
	}

	ProducerConsumerUtil(items)

	return nil

}
