package service

import (
	"time"

	"github.com/shivamk2406/item-inventory/database"
)

const (
	Host                  = "127.0.0.1:3306"
	User                  = "alpha"
	Password              = "alpha"
	Name                  = "inventory"
	MaxIdleConnections    = 10
	MaxOpenConnections    = 10
	MaxConnectionLifeTime = time.Minute * 3
	MaxConnectionIdleTime = time.Minute * 3
	DisableTLS            = true
)

func LoadAppConfig() database.Config {
	return database.Config{
		User:                  User,
		Password:              Password,
		Host:                  Host,
		Name:                  Name,
		MaxIdleConnections:    MaxIdleConnections,
		MaxOpenConnections:    MaxOpenConnections,
		MaxConnectionLifeTime: MaxConnectionLifeTime,
		MaxConnectionIdleTime: MaxConnectionIdleTime,
		DisableTLS:            DisableTLS}
}
