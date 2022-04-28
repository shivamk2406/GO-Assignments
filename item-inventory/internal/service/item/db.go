package item

import (
	"log"

	"gorm.io/gorm"
)

type DB interface {
	GetInventoryItem() ([]Item, error)
	BatchInsertion([]Item)
}

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (*Repository, error) {
	return &Repository{db: db}, nil
}

func (r *Repository) GetInventoryItem() ([]Item, error) {
	var items []Item
	if err := r.db.Find(&items).Error; err != nil {
		log.Println(err)
		return []Item{}, err
	}

	return items, nil
}

func (r *Repository) BatchInsertion(items []Item) {
	r.db.Create(&items)
}
