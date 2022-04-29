package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/shivamk2406/item-inventory/internal/config"
	"github.com/shivamk2406/item-inventory/internal/service/item"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	User                  string        `json:"user,omitempty"`
	Password              string        `json:"password,omitempty"`
	Host                  string        `json:"host,omitempty"`
	Name                  string        `json:"name,omitempty"`
	MaxIdleConnections    int           `json:"max_idle_connections,omitempty"`
	MaxOpenConnections    int           `json:"max_open_connections,omitempty"`
	MaxConnectionLifeTime time.Duration `json:"max_connection_life_time,omitempty"`
	MaxConnectionIdleTime time.Duration `json:"max_connection_idle_time,omitempty"`
	DisableTLS            bool          `json:"disable_tls,omitempty"`
}

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

	if !db.Migrator().HasTable(&item.Item{}) {
		err := db.AutoMigrate(&item.Item{})
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

func InitializeDB(conf config.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var dbOnce sync.Once
	var err error
	dbOnce.Do(func() {

		db, _, err = Open(conf)
	})
	return db, err
}
