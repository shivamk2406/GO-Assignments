package e2e

import (
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/subscriptions"
	"gorm.io/gorm"
)

func InsertGenreData(db *gorm.DB) {
	var news = []news.News{{NewsID: 1, GenreID: 1, Description: "Oppo"}}
	var genres = []subscriptions.Genre{{ID: 1, Name: "Tech", Description: "Get Latest Tech News"}}

	db.Create(&news)
	db.Create(&genres)
}
