package app

import (
	"fmt"

	"github.com/shivamk2406/item-inventory/internal/service"
)

func Start() {
	repo, err := InitializeFactory()
	if err != nil {
		fmt.Println(err)
	}
	service.ProcessorBuffered(repo)
}
