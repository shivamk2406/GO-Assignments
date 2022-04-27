package consumer

import (
	"sync"

	"github.com/shivamk2406/item-inventory/internal/service/item"
)

func Consumer(c chan item.Item, invoices *[]item.Invoice, wg *sync.WaitGroup, mutex *sync.Mutex) {
	for val := range c {
		mutex.Lock()
		*invoices = append(*invoices, val.ItemInvoice())
		mutex.Unlock()
	}

	wg.Done()

}
