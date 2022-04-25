package service

import (
	"log"

	"github.com/shivamk2406/item-inventory/database"
	"github.com/shivamk2406/item-inventory/domain/item"
	"gorm.io/gorm"
)

type DB interface {
	GetInventoryItem() ([]item.Item, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo() (error, *repository) {

	db, err := database.Open(database.Config{
		User:                  User,
		Password:              Password,
		Host:                  Host,
		Name:                  Name,
		MaxIdleConnections:    MaxIdleConnections,
		MaxOpenConnections:    MaxOpenConnections,
		MaxConnectionLifeTime: MaxConnectionLifeTime,
		MaxConnectionIdleTime: MaxConnectionIdleTime,
		DisableTLS:            DisableTLS})

	if err != nil {
		log.Println(err)
		return err, &repository{}
	}

	return nil, &repository{db: db}
}

func (r *repository) GetInventoryItem() ([]item.Item, error) {

	var items []item.Item

	if err := r.db.Find(&items).Error; err != nil {
		log.Println(err)
		return []item.Item{}, err
	}

	return items, nil
}
