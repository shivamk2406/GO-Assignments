package service

import "time"

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
