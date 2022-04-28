package app

import (
	"fmt"

	"github.com/shivamk2406/item-inventory/internal/config"
	"github.com/shivamk2406/item-inventory/internal/service"
)

func Start() {

	config, err := config.LoadDatabaseConfig()
	if err != nil {
		fmt.Println(err)
	}

	repo := Wire(config)

	service.Util(repo)
	// err := service.Init()
	// if err != nil {
	// 	log.Println(err)
	// }
}
