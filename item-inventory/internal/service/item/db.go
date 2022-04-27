package item

import (
	"log"

	"gorm.io/gorm"
)

type DB interface {
	GetInventoryItem() ([]Item, error)
	BatchInsertion([]Item)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (*repository, error) {
	return &repository{db: db}, nil
}

func (r *repository) GetInventoryItem() ([]Item, error) {
	var items []Item
	if err := r.db.Find(&items).Error; err != nil {
		log.Println(err)
		return []Item{}, err
	}

	return items, nil
}

func (r *repository) BatchInsertion(items []Item) {
	r.db.Create(&items)
}
