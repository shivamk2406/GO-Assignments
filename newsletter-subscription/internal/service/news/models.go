package news

type News struct {
	NewsID      int    `gorm:"primary_key;column:newsid;type:bigint;"`
	GenreID     int    `gorm:"column:genreid;type:bigint;"`
	Heading     string `gorm:"column:heading;"`
	Description string `gorm:"column:description;"`
}
