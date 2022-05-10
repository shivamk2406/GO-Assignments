package e2e

import (
	"github.com/shivamk2406/newsletter-subscriptions/internal/models"
	"gorm.io/gorm"
)

func InsertGenreData(db *gorm.DB) {
	var news = []models.News{{NewsID: 1, GenreID: 1, Description: "Oppo"}}
	var genres = []models.Genre{{ID: 1, Name: "Tech", Description: "Get Latest Tech News"}}

	db.Create(&news)
	db.Create(&genres)
}
