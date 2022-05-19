package database

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/subscriptions"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open(cfg config.Config) (*gorm.DB, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=10s&charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, errors.Wrap(err, "database: could not set sql.DB params")
	}

	if !db.Migrator().HasTable(&users.User{}) {
		err := db.AutoMigrate(&users.User{})
		if err != nil {
			return db, nil, err
		}
	}

	if !db.Migrator().HasTable(&news.News{}) {
		err := db.AutoMigrate(&news.News{})
		if err != nil {
			return db, nil, err
		}
	}
	if !db.Migrator().HasTable(&subscriptions.Genre{}) {
		err := db.AutoMigrate(&subscriptions.Genre{})
		if err != nil {
			return db, nil, err
		}
	}
	if !db.Migrator().HasTable(&subscriptions.Subscriptions{}) {
		err := db.AutoMigrate(&subscriptions.Subscriptions{})
		if err != nil {
			return db, nil, err
		}
	}
	if !db.Migrator().HasTable(&subscriptions.SubscriptionGenre{}) {
		err := db.AutoMigrate(&subscriptions.SubscriptionGenre{})
		if err != nil {
			return db, nil, err
		}
	}
	sqlDB.SetConnMaxIdleTime(cfg.Database.MaxConnectionIdleTime)
	sqlDB.SetConnMaxLifetime(cfg.Database.MaxConnectionLifeTime)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConnections)

	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("failed to close db connections %v", err)
		}
	}
	return db, cleanup, nil
}
