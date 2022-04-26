package service

import (
	"log"

	"github.com/shivamk2406/item-inventory/domain/item"
	"gorm.io/gorm"
)

type DB interface {
	GetInventoryItem() ([]item.Item, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (error, *repository) {
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
