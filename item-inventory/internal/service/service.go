package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/shivamk2406/item-inventory/internal/config"
	"github.com/shivamk2406/item-inventory/internal/service/consumer"
	"github.com/shivamk2406/item-inventory/internal/service/item"
	producer "github.com/shivamk2406/item-inventory/internal/service/producer"
	"github.com/shivamk2406/item-inventory/pkg/database"
	"gorm.io/gorm"
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
	consumerCount, producerCount, err := config.LoadRoutineConfig()
	if err != nil {
		fmt.Println(err)
	}

	c := make(chan item.Item)
	var invoices []item.Invoice
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(producerCount + consumerCount)

	go producer.Producer(repo, c, &wg)
	for i := 0; i < consumerCount; i++ {
		go consumer.Consumer(c, &invoices, &wg, &mutex)
	}

	wg.Wait()
	fmt.Printf("Total Length of invoice generated %d \n", len(invoices))
}

func ProviderDB(conf config.Config) *gorm.DB {
	var db *gorm.DB
	var dbOnce sync.Once

	dbOnce.Do(func() { db, _, _ = database.Open(conf) })
	return db
}
