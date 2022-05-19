package subscriptions

type Subscriptions struct {
	ID      int    `gorm:"primary_key;column:id;"`
	Name    string `gorm:"column:name;"`
	Renewal int    `gorm:"column:renewal;"`
	Price   int    `gorm:"column:price;"`
}

type Genre struct {
	ID          int    `gorm:"primary_key;column:id;"`
	Name        string `gorm:"column:name;"`
	Description string `gorm:"column:description;"`
}

type SubscriptionGenre struct {
	SubsID int `gorm:"primary_key;column:subscriptions_id;type:bigint;"`
	GenID  int `gorm:"primary_key;column:genre_id;type:bigint;"`
}
