package service

import (
	"github.com/shivamk2406/item-inventory/domain/item"
)

const (
	RoutineCount   = 5
	BufferCapacity = 1
)

func getItemFromDB(repo *repository) (error, []item.Item) {
	items, err := repo.GetInventoryItem()
	if err != nil {
		return err, []item.Item{}
	}

	return nil, items
}

func ProducerConsumerUtil(repo *repository) (error, []item.Invoice) {
	err, items := getItemFromDB(repo)
	if err != nil {
		return err, []item.Invoice{}
	}

	consumerChannel := make(chan item.Item, BufferCapacity)
	producerChannel := make(chan item.Invoice, BufferCapacity)
	var itemInvoices []item.Invoice

	for i := 0; i < RoutineCount; i++ {
		go worker(consumerChannel, producerChannel)
	}

	for i := 0; i < len(items); i++ {
		consumerChannel <- items[i]
	}
	close(consumerChannel)

	for j := 0; j < len(items); j++ {
		itemInvoices = append(itemInvoices, <-producerChannel)
	}
	return nil, itemInvoices
}

func worker(consumer <-chan item.Item, producer chan<- item.Invoice) {
	for val := range consumer {
		producer <- val.ItemInvoice()
	}
}

func ProducerConsumerUtil1(repo *repository) (error, []item.Invoice) {
	err, items := getItemFromDB(repo)
	if err != nil {
		return err, []item.Invoice{}
	}

	consumerChannel := make(chan item.Item)
	var itemInvoices []item.Invoice

	go consumer(consumerChannel, items)

	for n := range consumerChannel {
		itemInvoices = append(itemInvoices, n.ItemInvoice())
	}
	return nil, itemInvoices
}

func consumer(consumerChannel chan item.Item, items []item.Item) {
	for i := 0; i < len(items); i++ {
		consumerChannel <- items[i]
	}
	close(consumerChannel)
}
