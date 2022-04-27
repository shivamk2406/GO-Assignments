package service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/shivamk2406/item-inventory/internal/config"
	"github.com/shivamk2406/item-inventory/internal/consumer"
	producer "github.com/shivamk2406/item-inventory/internal/producer"
	"github.com/shivamk2406/item-inventory/internal/service/item"
	"github.com/shivamk2406/item-inventory/internal/service/item/enum"
	"github.com/shivamk2406/item-inventory/pkg/database"
)

func Init() error {
	config, err := config.LoadDatabaseConfig()
	if err != nil {
		return err
	}

	db, cleanup, err := database.Open(config)
	if err != nil {
		return err
	}
	defer cleanup()

	repo, err := item.NewRepo(db)
	if err != nil {
		return err
	}
	start := time.Now()
	//BatchInsertion(repo)

	Util(repo)
	fmt.Println(time.Since(start))
	return nil

}

func Util(repo item.DB) {

	routineCount, err := config.LoadRoutineConfig()
	if err != nil {
		fmt.Println(err)
	}

	c := make(chan item.Item)
	var invoices []item.Invoice
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(routineCount)

	go producer.Producer(repo, c, &wg)
	for i := 0; i < routineCount-1; i++ {
		go consumer.Consumer(c, &invoices, &wg, &mutex)
	}

	wg.Wait()
	fmt.Printf("Total Length of invoice generated %d \n", len(invoices))
}

func BatchInsertion(repo item.DB) {
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

	repo.BatchInsertion(items)

	fmt.Println(len(items))

}
