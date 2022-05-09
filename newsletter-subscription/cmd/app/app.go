package app

import (
	"log"

	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	"github.com/shivamk2406/newsletter-subscriptions/internal/pkg/database"
)

func Start() {
	conf, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Println(err)
	}

	db, cleanup, err := database.Open(conf)
	if err != nil {
		log.Println(err)
	}

	print(db)
	defer cleanup()

}
