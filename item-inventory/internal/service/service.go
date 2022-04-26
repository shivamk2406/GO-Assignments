package service

import (
	"fmt"

	"github.com/shivamk2406/item-inventory/database"
	"github.com/shivamk2406/item-inventory/domain/item"
)

func Init() error {
	err, config := LoadAppConfig()
	if err != nil {
		return err
	}

	db, err := database.Open(config)
	if err != nil {
		return err
	}

	err, repo := NewRepo(db)
	if err != nil {
		return err
	}

	err, itemInvoices := ProducerConsumerUtil(repo)
	if err != nil {
		return err
	}

	displayInvoices(itemInvoices)

	return nil

}

func displayInvoices(invoices []item.Invoice) {
	for _, val := range invoices {
		fmt.Printf("%v \n", val)
	}
}
