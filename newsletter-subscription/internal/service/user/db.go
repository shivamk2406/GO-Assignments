package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type CreateUserRequest struct {
	Name  string
	Email string
}

type AuthenticateUserRequest struct {
	Email string
}

type NewUser struct {
	Name   string
	Email  string
	Active bool
}

type AuthenticateUserResponse struct {
	IsAuthenticated bool
	User            NewUser
}

type ListActiveUsers struct{}

type ListActiveUsersResponse struct {
	ActiveUsers []NewUser
}

type UsersDB interface {
	createUser(ctx context.Context, in CreateUserRequest) (NewUser, error)
	authenticateUser(ctx context.Context, in AuthenticateUserRequest) (AuthenticateUserResponse, error)
	listActiveUsers(ctx context.Context, in ListActiveUsers) (ListActiveUsersResponse, error)
}
type Repository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UsersDB {
	return &Repository{db: db}
}

func (r Repository) createUser(ctx context.Context, in CreateUserRequest) (NewUser, error) {
	log.Printf("Received: %v %v", in.Name, in.Email)
	user := User{ID: 1, Email: in.Email, Name: in.Name, Active: false, StartDate: time.Now(), EndDate: time.Now()}
	if err := r.db.Create(&user).Error; err != nil {
		return NewUser{}, err
	}

	log.Println(user)
	return NewUser{Name: in.Name, Email: in.Email, Active: false}, nil
}

func (r Repository) authenticateUser(ctx context.Context, in AuthenticateUserRequest) (AuthenticateUserResponse, error) {

	var user User
	log.Printf("Received : %v", in.Email)
	log.Printf("Inside DB: %s", in)
	if err := r.db.First(&user, "email = ?", in.Email).Error; err != nil {
		return AuthenticateUserResponse{IsAuthenticated: false}, err
	}
	return AuthenticateUserResponse{IsAuthenticated: true, User: NewUser{Name: user.Name, Email: user.Email, Active: user.Active}}, nil
}

func (r Repository) listActiveUsers(ctx context.Context, in ListActiveUsers) (ListActiveUsersResponse, error) {
	var activeUsersDB []User
	r.db.Where("end_time > ?", time.Now()).Find(&activeUsersDB)
	var activeUsers []NewUser

	for _, val := range activeUsersDB {
		activeUsers = append(activeUsers, NewUser{Name: val.Name, Email: val.Email, Active: val.Active})
	}

	var listActiveUser ListActiveUsersResponse
	listActiveUser.ActiveUsers = activeUsers
	fmt.Println(listActiveUser)
	return listActiveUser, nil
}
