package database

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	"github.com/shivamk2406/newsletter-subscriptions/internal/models"
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

	if !db.Migrator().HasTable(&models.User{}) {
		err := db.AutoMigrate(&models.User{})
		if err != nil {
			return db, nil, err
		}
	}
	if !db.Migrator().HasTable(&models.News{}) {

		err := db.AutoMigrate(&models.News{})
		if err != nil {
			return db, nil, err
		}
	}
	if !db.Migrator().HasTable(&models.Genre{}) {
		err := db.AutoMigrate(&models.Genre{})
		if err != nil {
			return db, nil, err
		}
	}
	if !db.Migrator().HasTable(&models.Subscriptions{}) {
		err := db.AutoMigrate(&models.Subscriptions{})
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
