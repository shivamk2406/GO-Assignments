package subscriptions

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shivamk2406/newsletter-subscriptions/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type GetPlansRequests struct {
}

type Genre struct {
	Name string
}
type Subscription struct {
	Id       int32
	Name     string
	Validity int32
	Genres   []*Genre
}
type Plans struct {
	Subs []*Subscription
}

type SubscriptionRequest struct {
	Email string
}

type SetSubscriptionRequest struct {
	Email  string
	Subsid int32
}

type SetSubscriptionResponse struct {
	Email     string
	Active    bool
	Starttime timestamppb.Timestamp
	Validity  int32
}

type SubscriptionDB interface {
	getPlans(ctx context.Context, in GetPlansRequests) (Plans, error)
	getSubscription(ctx context.Context, in SubscriptionRequest) (Subscription, error)
	setSubscription(ctx context.Context, in SetSubscriptionRequest) (SetSubscriptionResponse, error)
}

type Repository struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) SubscriptionDB {
	return &Repository{db: db}
}

func (r Repository) getGenresById(id int) []*Genre {
	var subscriptionGenres []models.SubscriptionGenre
	r.db.Where("subscriptions_id = ?", id).Find(&subscriptionGenres)

	var genreIds []int
	for _, val := range subscriptionGenres {
		genreIds = append(genreIds, val.GenID)
	}

	var genres []models.Genre
	for _, val := range genreIds {
		var genre models.Genre
		r.db.Where("ID = ?", val).Find(&genre)
		genres = append(genres, genre)
	}

	var genresField []*Genre
	for _, val := range genres {
		genre := Genre{Name: val.Name}
		genresField = append(genresField, &genre)
	}
	return genresField
}

func (r *Repository) getPlans(ctx context.Context, in GetPlansRequests) (Plans, error) {
	var subs []models.Subscriptions
	if err := r.db.Find(&subs).Error; err != nil {
		return Plans{}, err
	}

	var plans []*Subscription
	for _, val := range subs {
		genres := r.getGenresById(val.ID)
		plans = append(plans, &Subscription{Name: val.Name, Id: int32(val.ID), Validity: int32(val.Renewal), Genres: genres})
	}
	return Plans{Subs: plans}, nil
}

func (r Repository) getSubscription(ctx context.Context, in SubscriptionRequest) (Subscription, error) {
	log.Printf("Received : %v", in.Email)
	var user models.User
	row := r.db.First(&user, "email = ?", in.Email)
	if err := row.Error; err != nil {
		log.Println(err)
		return Subscription{}, fmt.Errorf("You have not Subscribed to any plans")
	}

	var subs models.Subscriptions
	if err := r.db.Where("ID = ?", user.SubsID).Find(&subs).Error; err != nil {
		return Subscription{}, err
	}

	genresField := r.getGenresById(subs.ID)
	return Subscription{Id: int32(subs.ID), Name: subs.Name, Validity: int32(subs.Renewal), Genres: genresField}, nil
}

func (r Repository) setSubscription(ctx context.Context, in SetSubscriptionRequest) (SetSubscriptionResponse, error) {
	log.Printf("Recieved Email %s and Subsid %d ", in.Email, in.Subsid)
	var subs models.Subscriptions
	r.db.Where("ID = ?", in.Subsid).Find(&subs)
	// Update with conditions
	if err := r.db.Model(&models.User{}).Where("email = ?", in.Email).
		Updates(map[string]interface{}{
			"active":     true,
			"start_time": time.Now(),
			"subsid":     in.Subsid,
			"validity":   subs.Renewal}).Error; err != nil {
		return SetSubscriptionResponse{}, err
	}
	return SetSubscriptionResponse{Email: in.Email, Active: true, Starttime: *timestamppb.Now(), Validity: int32(subs.Renewal)}, nil
}
