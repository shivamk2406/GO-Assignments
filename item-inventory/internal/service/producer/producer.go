package producer

import (
	"fmt"
	"sync"

	"github.com/shivamk2406/item-inventory/internal/service/item"
)

func Producer(repo item.DB, c chan item.Item, wg *sync.WaitGroup) {
	items, err := repo.GetInventoryItem()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range items {
		c <- val
	}
	close(c)
	wg.Done()
}
