package app

import (
	"github.com/shivamk2406/item-inventory/internal/service"
)

func Start() {
	repo := Wire()
	service.RunApp(repo)
}
