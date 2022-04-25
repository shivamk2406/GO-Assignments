package database

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
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

func Open(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "database: could not set sql.DB params")
	}

	sqlDB.SetConnMaxIdleTime(cfg.MaxConnectionIdleTime)
	sqlDB.SetConnMaxLifetime(cfg.MaxConnectionLifeTime)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)

	return db, nil

}
