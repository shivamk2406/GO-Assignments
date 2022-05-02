package e2e

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shivamk2406/item-inventory/internal/service/item"
	"github.com/shivamk2406/item-inventory/internal/service/item/enum"
)

func DataGenerator(repo item.DB) {
	rand.Seed(time.Now().UnixNano())
	var items []item.Item
	names := [6]string{
		"pen",
		"paper",
		"Book",
		"Copy",
		"Pencil",
		"Eraser",
	}

	types := [3]string{
		"manufactured",
		"imported",
		"raw",
	}

	for i := 0; i < 10000; i++ {
		types, _ := enum.ItemTypeString(types[rand.Intn(2)+1])
		item := item.Item{
			Name:     names[rand.Intn(5)+1],
			Price:    100 + rand.Float64()*100,
			Quantity: rand.Intn(8) + 2,
			Type:     types,
		}
		fmt.Println(item)
		items = append(items, item)
	}

	//repo.BatchInsertion(items)

	fmt.Println(len(items))

}
