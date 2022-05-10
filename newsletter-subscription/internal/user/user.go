package user

import (
	"context"
	"fmt"
	"log"

	"github.com/shivamk2406/newsletter-subscriptions/internal/models"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type UserDB interface {
	GetSubscription(string)
	GetNews(string)
}

type userManagementServer struct {
	pb.UnimplementedUserManagementServer
	db *gorm.DB
}

func NewRepo(db *gorm.DB) pb.UserManagementServer {
	return &userManagementServer{db: db}
}

func (r *userManagementServer) GetGenres() ([]models.Genre, error) {
	var genres []models.Genre
	if err := r.db.Find(&genres).Error; err != nil {
		log.Println(err)
		return []models.Genre{}, err
	}
	return genres, nil
}

func (r *userManagementServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserReponse, error) {
	log.Printf("Received: %v", in.Subsid)
	log.Printf("Created User: %s %s %d %d %d %v", in.Name, in.Email, 1, 30, in.Subsid, timestamppb.Now().AsTime())
	user := models.User{ID: 1, Email: in.Email, Name: in.Name, StartDate: timestamppb.Now().AsTime(), SubsID: int(in.Subsid), Active: true, Validity: 30}
	r.db.Create(&user)
	return &pb.CreateUserReponse{Name: in.Name, Email: in.Email, Active: 1, Validity: 30, Subsid: in.Subsid, Starttime: timestamppb.Now()}, nil
}

func (r *userManagementServer) GetAllSubscriptions() ([]models.Subscriptions, error) {
	var subscriptions []models.Subscriptions
	if err := r.db.Find(&subscriptions).Error; err != nil {
		log.Println(err)
		return []models.Subscriptions{}, err
	}
	return subscriptions, nil
}

func (r *userManagementServer) AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserRespone, error) {
	var user models.User
	log.Printf("Received : %v", in.Email)
	r.db.First(&user, "email = ?", in.Email)
	if user == (models.User{}) {
		return &pb.AuthenticateUserRespone{IsAuthenticated: false}, fmt.Errorf("user with %s email do not exists please sign up", in.Email)
	}
	fmt.Printf("Welcome %s", user.Name)
	return &pb.AuthenticateUserRespone{IsAuthenticated: true}, nil
}

func (r *userManagementServer) SetSubscription(subsid int, email string) error {
	return nil
}
