package news

import (
	"context"
	"log"

	"github.com/shivamk2406/newsletter-subscriptions/internal/models"
	"gorm.io/gorm"
)

type SingleNews struct {
	heading     string
	description string
}

type News struct {
	Newss []SingleNews
}

type GetNewsByGenreRequest struct {
	Genre string
}

type GetNewsRequest struct {
	Subsid int32
}

type NewsDB interface {
	getNewsByGenre(ctx context.Context, in GetNewsByGenreRequest) (News, error)
	getNews(ctx context.Context, in GetNewsRequest) (News, error)
}

type Repository struct {
	db *gorm.DB
}

func NewNewsRepo(db *gorm.DB) NewsDB {
	return &Repository{db: db}
}

func (r Repository) getNews(ctx context.Context, in GetNewsRequest) (News, error) {
	log.Printf("Received %v", in.Subsid)
	var subscriptionGenres []models.SubscriptionGenre
	if err := r.db.Where("subscriptions_id = ?", in.Subsid).Find(&subscriptionGenres).Error; err != nil {
		return News{}, err
	}

	var genreIds []int
	for _, val := range subscriptionGenres {
		genreIds = append(genreIds, val.GenID)
	}

	var newsCollection []models.News
	if err := r.db.Where("genreid IN ?", genreIds).Find(&newsCollection).Error; err != nil {
		return News{}, nil
	}

	var newsString []SingleNews
	for _, val := range newsCollection {
		newsString = append(newsString, SingleNews{heading: val.Heading, description: val.Description})
	}

	return News{Newss: newsString}, nil
}

func (r Repository) getNewsByGenre(ctx context.Context, in GetNewsByGenreRequest) (News, error) {
	log.Printf("Received %v", in.Genre)
	var genre models.Genre

	if err := r.db.Where("name = ?", in.Genre).Find(&genre).Error; err != nil {
		return News{}, err
	}

	var news []models.News
	if err := r.db.Where("genreid = ?", genre.ID).Find(&news).Error; err != nil {
		return News{}, err
	}

	var newsString []SingleNews
	for _, val := range news {
		newsString = append(newsString, SingleNews{heading: val.Heading, description: val.Description})
	}

	return News{Newss: newsString}, nil
}