package models

import (
	"time"
)

type User struct {
	ID            int           `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	Email         string        `gorm:"primary_key;column:email;"`
	Name          string        `gorm:"column:name;"`
	StartDate     time.Time     `gorm:"column:start_time;type:datetime;"`
	Subs_id       int           `gorm:"column:subsid;type:bigint;"`
	Active        bool          `gorm:"column:active;type:tinyint;default:1;"`
	Validity      int           `gorm:"column:validity;NOT NULL,type:ubigint;"`
	Subscriptions Subscriptions `gorm:"foreignKey:id"`
}

type Subscriptions struct {
	ID      int     `gorm:"primary_key;column:id;"`
	Name    string  `gorm:"column:name;"`
	Renewal int     `gorm:"column:renewal;"`
	Genres  []Genre `gorm:"many2many:subscription_genre;"`
}

type Genre struct {
	ID          int    `gorm:"primary_key;column:id;"`
	Name        string `gorm:"column:name;"`
	Description string `gorm:"column:description;"`
	News        []News `gorm:"foreignKey:GenreID;"`
}

type News struct {
	NewsID      int    `gorm:"primary_key;column:newsid;type:bigint;"`
	GenreID     int    `gorm:"column:genreid;type:bigint;"`
	Description string `gorm:"column:description;"`
}
