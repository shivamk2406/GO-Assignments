package app

import (
	"fmt"
	"log"

	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	"github.com/shivamk2406/newsletter-subscriptions/internal/pkg/database"
	"github.com/shivamk2406/newsletter-subscriptions/internal/user"
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

	repo := user.NewRepo(db)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(repo)
	// genres, err := repo.GetGenres()
	// if err != nil {
	// 	log.Println(err)
	// }
	//fmt.Println(genres)
	//subs, err := repo.GetAllSubscriptions()
	//if err != nil {
	//log.Println(err)
	//}
	//for _, val := range subs {
	//fmt.Printf("%d %s %d %d \n", val.ID, val.Name, val.Price, val.Renewal)
	//}
	// users, err := repo.AuthenticateUser("test@test.com")
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Printf("User Info:%s %s %d %d  ", users.Email, users.Name, users.ID, users.SubsID)
	//e2e.InsertGenreData(db)
	defer cleanup()

}
