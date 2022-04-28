package service

import (
	"fmt"
	"log"
	"sync"

	"github.com/shivamk2406/item-inventory/internal/config"
	"github.com/shivamk2406/item-inventory/internal/service/consumer"
	"github.com/shivamk2406/item-inventory/internal/service/item"
	producer "github.com/shivamk2406/item-inventory/internal/service/producer"
	"github.com/shivamk2406/item-inventory/pkg/database"
	"gorm.io/gorm"
)

func RunApp(repo item.DB) {
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

	dbOnce.Do(func() {
		var err error
		db, _, err = database.Open(conf)
		if err != nil {
			log.Println(err)
		}
	})
	return db
}

func ProviderRepo(db *gorm.DB) *item.Repository {
	var repo *item.Repository
	var repoOnce sync.Once

	repoOnce.Do(func() {
		var err error
		repo, err = item.NewRepo(db)
		if err != nil {
			log.Println(err)
		}
	})
	return repo
}

func ProviderConfig() config.Config {
	var conf config.Config
	var confOnce sync.Once

	confOnce.Do(func() {
		var err error
		conf, err = config.LoadDatabaseConfig()
		if err != nil {
			log.Println(err)
		}
	})
	return conf
}
