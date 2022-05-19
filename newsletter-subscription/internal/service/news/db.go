package news

import (
	"context"
	"log"

	"github.com/shivamk2406/newsletter-subscriptions/internal/service/subscriptions"
	"gorm.io/gorm"
)

type SingleNews struct {
	heading     string
	description string
}

type news struct {
	Newss []SingleNews
}

type GetNewsByGenreRequest struct {
	Genre string
}

type GetNewsRequest struct {
	Subsid int32
}

type NewsDB interface {
	getNewsByGenre(ctx context.Context, in GetNewsByGenreRequest) (news, error)
	getNews(ctx context.Context, in GetNewsRequest) (news, error)
}

type Repository struct {
	db *gorm.DB
}

func NewNewsRepo(db *gorm.DB) NewsDB {
	return &Repository{db: db}
}

func (r Repository) getNews(ctx context.Context, in GetNewsRequest) (news, error) {
	log.Printf("Received %v", in.Subsid)
	var subscriptionGenres []subscriptions.SubscriptionGenre
	if err := r.db.Where("subscriptions_id = ?", in.Subsid).Find(&subscriptionGenres).Error; err != nil {
		return news{}, err
	}

	var genreIds []int
	for _, val := range subscriptionGenres {
		genreIds = append(genreIds, val.GenID)
	}

	var newsCollection []News
	if err := r.db.Where("genreid IN ?", genreIds).Find(&newsCollection).Error; err != nil {
		return news{}, nil
	}

	var newsString []SingleNews
	for _, val := range newsCollection {
		newsString = append(newsString, SingleNews{heading: val.Heading, description: val.Description})
	}

	return news{Newss: newsString}, nil
}

func (r Repository) getNewsByGenre(ctx context.Context, in GetNewsByGenreRequest) (news, error) {
	log.Printf("Received %v", in.Genre)
	var genre subscriptions.Genre

	if err := r.db.Where("name = ?", in.Genre).Find(&genre).Error; err != nil {
		return news{}, err
	}

	var newsList []News
	if err := r.db.Where("genreid = ?", genre.ID).Find(&newsList).Error; err != nil {
		return news{}, err
	}

	var newsString []SingleNews
	for _, val := range newsList {
		newsString = append(newsString, SingleNews{heading: val.Heading, description: val.Description})
	}

	return news{Newss: newsString}, nil
}
