package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shivamk2406/newsletter-subscriptions/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type CreateUserRequest struct {
	Name  string
	Email string
}

type GetPlansRequests struct {
}

type AuthenticateUserRequest struct {
	Email string
}

type SubscriptionRequest struct {
	Email string
}

type SetSubscriptionRequest struct {
	Email  string
	Subsid int32
}

type GetNewsByGenreRequest struct {
	Genre string
}

type GetNewsRequest struct {
	Subsid int32
}

type NewUser struct {
	Name   string
	Email  string
	Active bool
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

type AuthenticateUserResponse struct {
	IsAuthenticated bool
	User            NewUser
}

type SetSubscriptionResponse struct {
	Email     string
	Active    bool
	Starttime timestamppb.Timestamp
	Validity  int32
}

type News struct {
	Newss []string
}

type DB interface {
	CreateUser(ctx context.Context, in CreateUserRequest) (NewUser, error)
	GetGenresById(id int) []*Genre
	GetPlans(ctx context.Context, in GetPlansRequests) (Plans, error)
	GetAllSubscriptions() ([]models.Subscriptions, error)
	AuthenticateUser(ctx context.Context, in AuthenticateUserRequest) (AuthenticateUserResponse, error)
	GetSubscription(ctx context.Context, in SubscriptionRequest) (Subscription, error)
	SetSubsciption(ctx context.Context, in SetSubscriptionRequest) (SetSubscriptionResponse, error)
	GetNewsByGenre(ctx context.Context, in GetNewsByGenreRequest) (News, error)
	GetNews(ctx context.Context, in GetNewsRequest) (News, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetGenresById(id int) []*Genre {
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

func (r Repository) CreateUser(ctx context.Context, in CreateUserRequest) (NewUser, error) {
	log.Printf("Received: %v %v", in.Name, in.Email)
	user := models.User{ID: 1, Email: in.Email, Name: in.Name, Active: false, StartDate: time.Now()}
	if err := r.db.Create(&user).Error; err != nil {
		return NewUser{}, err
	}

	log.Println(user)
	return NewUser{Name: in.Name, Email: in.Email, Active: false}, nil
}

func (r Repository) GetAllSubscriptions() ([]models.Subscriptions, error) {
	var subscriptions []models.Subscriptions
	if err := r.db.Find(&subscriptions).Error; err != nil {
		log.Println(err)
		return []models.Subscriptions{}, err
	}

	return subscriptions, nil
}

func (r Repository) GetPlans(ctx context.Context, in GetPlansRequests) (Plans, error) {
	var subs []models.Subscriptions
	if err := r.db.Find(&subs).Error; err != nil {
		return Plans{}, err
	}

	var plans []*Subscription
	for _, val := range subs {
		genres := r.GetGenresById(val.ID)
		plans = append(plans, &Subscription{Name: val.Name, Id: int32(val.ID), Validity: int32(val.Renewal), Genres: genres})
	}
	return Plans{Subs: plans}, nil

}
func (r Repository) AuthenticateUser(ctx context.Context, in AuthenticateUserRequest) (AuthenticateUserResponse, error) {
	var user models.User
	log.Printf("Received : %v", in.Email)
	log.Printf("Inside DB: %s", in)
	if err := r.db.First(&user, "email = ?", in.Email).Error; err != nil {
		return AuthenticateUserResponse{IsAuthenticated: false}, err
	}
	return AuthenticateUserResponse{IsAuthenticated: true, User: NewUser{Name: user.Name, Email: user.Email, Active: user.Active}}, nil
}

func (r Repository) GetSubscription(ctx context.Context, in SubscriptionRequest) (Subscription, error) {
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

	genresField := r.GetGenresById(subs.ID)
	return Subscription{Id: int32(subs.ID), Name: subs.Name, Validity: int32(subs.Renewal), Genres: genresField}, nil
}

func (r Repository) SetSubsciption(ctx context.Context, in SetSubscriptionRequest) (SetSubscriptionResponse, error) {
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

func (r Repository) GetNewsByGenre(ctx context.Context, in GetNewsByGenreRequest) (News, error) {
	log.Printf("Received %v", in.Genre)
	var genre models.Genre

	if err := r.db.Where("name = ?", in.Genre).Find(&genre).Error; err != nil {
		return News{}, err
	}

	var news []models.News
	if err := r.db.Where("genreid = ?", genre.ID).Find(&news).Error; err != nil {
		return News{}, err
	}

	var newsString []string
	for _, val := range news {
		newsString = append(newsString, val.Description)
	}

	return News{Newss: newsString}, nil
}

func (r Repository) GetNews(ctx context.Context, in GetNewsRequest) (News, error) {
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

	var newsString []string
	for _, val := range newsCollection {
		newsString = append(newsString, val.Description)
	}

	return News{Newss: newsString}, nil
}
