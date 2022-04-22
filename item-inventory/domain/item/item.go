package item

import "github.com/shivamk2406/item-inventory/domain/item/enum"

type Item struct {
	Name     string
	Price    float64
	Quantity int
	Type     enum.ItemType
}
