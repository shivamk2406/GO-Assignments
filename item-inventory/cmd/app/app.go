package app

import (
	"fmt"

	"github.com/shivamk2406/item-inventory/internal/service"
)

func Start() {
	repo, err := InitializeEvent()
	if err != nil {
		fmt.Println(err)
	}
	service.ProducerConsumerUtil(repo)
}
