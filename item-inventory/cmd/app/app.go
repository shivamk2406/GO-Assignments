package app

import (
	"github.com/shivamk2406/item-inventory/internal/service"
)

func Start() {
	repo := InitializeEvent()
	service.ProducerConsumerUtil(repo)
}
