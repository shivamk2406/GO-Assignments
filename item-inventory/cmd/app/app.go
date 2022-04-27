package app

import (
	"log"

	"github.com/shivamk2406/item-inventory/internal/service"
)

func Start() {

	err := service.Init()
	if err != nil {
		log.Println(err)
	}
}
