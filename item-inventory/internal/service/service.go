package service

import (
	"fmt"
	"sync"

	"github.com/shivamk2406/item-inventory/internal/config"
	"github.com/shivamk2406/item-inventory/internal/service/consumer"
	"github.com/shivamk2406/item-inventory/internal/service/item"
	producer "github.com/shivamk2406/item-inventory/internal/service/producer"
)

func ProcessorUnbuffered(repo item.DB) {
	c := make(chan item.Item)
	var invoices []item.Invoice
	var wg sync.WaitGroup
	var mutex sync.Mutex

	consumerCount, err := config.LoadConsumerConfig()
	if err != nil {
		fmt.Println(err)
	}

	producerCount, err := config.LoadProducerConfig()
	if err != nil {
		fmt.Println(err)
	}

	wg.Add(producerCount + consumerCount)
	for i := 0; i < producerCount; i++ {
		go producer.Producer(repo, c, &wg)
	}

	for i := 0; i < consumerCount; i++ {
		go consumer.Consumer(c, &invoices, &wg, &mutex)
	}

	wg.Wait()
	fmt.Printf("Total Length of invoice generated %d \n", len(invoices))
}

func ProcessorBuffered(repo item.DB) {
	var invoices []item.Invoice
	var wg sync.WaitGroup
	var mutex sync.Mutex

	bufferCapacity, err := config.LoadChannelConfig()
	if err != nil {
		if err != nil {
			fmt.Println(err)
		}
	}

	c := make(chan item.Item, bufferCapacity)
	consumerCount, err := config.LoadConsumerConfig()
	if err != nil {
		fmt.Println(err)
	}

	producerCount, err := config.LoadProducerConfig()
	if err != nil {
		fmt.Println(err)
	}
	wg.Add(producerCount + consumerCount)
	for i := 0; i < producerCount; i++ {
		go producer.Producer(repo, c, &wg)
	}

	for i := 0; i < consumerCount; i++ {
		go consumer.Consumer(c, &invoices, &wg, &mutex)
	}

	wg.Wait()
	fmt.Printf("Total Length of invoice generated %d \n", len(invoices))

}
