package service

import (
	"fmt"
	"sync"

	"github.com/shivamk2406/item-inventory/internal/config"
	"github.com/shivamk2406/item-inventory/internal/consumer"
	producer "github.com/shivamk2406/item-inventory/internal/producer"
	"github.com/shivamk2406/item-inventory/internal/service/item"
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

	invoices := Util(repo)

	displayInvoices(invoices)
	return nil

}

func displayInvoices(invoices []item.Invoice) {
	for _, val := range invoices {
		fmt.Printf("%v \n", val)
	}
}

func Util(repo item.DB) []item.Invoice {
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

	return invoices
}
