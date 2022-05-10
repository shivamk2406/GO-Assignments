package server

import (
	"context"
	"log"

	grpcserver "github.com/shivamk2406/newsletter-subscriptions/internal/server/grpc"
	"github.com/shivamk2406/newsletter-subscriptions/internal/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RunServer() error {
	dsn := "alpha:alpha@tcp(127.0.0.1:3306)/news?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	context := context.Background()
	repo := user.NewRepo(db)

	return grpcserver.RunServer(context, repo)

}
