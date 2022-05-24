package users

import "time"

type User struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	Email     string    `gorm:"primary_key;column:email;"`
	Name      string    `gorm:"column:name;"`
	StartDate time.Time `gorm:"column:start_time;type:datetime;"`
	EndDate   time.Time `gorm:"column:end_time;type:datetime;"`
	SubsID    int       `gorm:"column:subsid;type:bigint;"`
	Active    bool      `gorm:"column:active;type:tinyint;default:0;"`
	Validity  int       `gorm:"column:validity;NOT NULL,type:ubigint;"`
}
