package service

import (
	"fmt"

	"github.com/shivamk2406/item-inventory/domain/item"
)

func ProducerConsumerUtil(items []item.Item) []item.Invoice {
	consumerChannel := make(chan item.Item, len(items))
	producerChannel := make(chan item.Invoice, len(items))
	var itemInvoices []item.Invoice

	go worker(consumerChannel, producerChannel)

	for i := 0; i < len(items); i++ {
		consumerChannel <- items[i]
	}
	close(consumerChannel)

	for j := 0; j < len(items); j++ {
		itemInvoices = append(itemInvoices, <-producerChannel)
		fmt.Println(itemInvoices[j])
	}
	return itemInvoices
}

func worker(consumer <-chan item.Item, producer chan<- item.Invoice) {
	for val := range consumer {
		producer <- val.ItemInvoice()

	}
}
