package main

import (
	"log"

	"github.com/shivamk2406/item-inventory/internal/service"
)

func main() {
	err := service.Init()
	if err != nil {
		log.Println(err)
	}
}
