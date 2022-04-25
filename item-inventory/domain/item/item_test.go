package item

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/shivamk2406/item-inventory/domain/item/enum"
)

type scenarioTestNewItem struct {
	description string
	name        string
	price       float64
	quantity    int
	itemType    string
	itemError   error
}

type scenarioTestEffectivePrice struct {
	description   string
	item          Item
	expectedPrice float64
	err           error
}

func TestNewItem(t *testing.T) {
	scenarios := []scenarioTestNewItem{
		{
			description: "All information is provided and Price is negative",
			name:        "Pen",
			price:       -87,
			quantity:    4,
			itemType:    "raw",
			itemError:   errors.Errorf("negative value"),
		},
		{
			description: "All correct information is provided",
			name:        "Book",
			price:       56,
			quantity:    4,
			itemType:    "manufactured",
			itemError:   nil,
		},
		{
			description: "All information is provided and quantity is negative",
			name:        "Pen",
			price:       87,
			quantity:    -4,
			itemType:    "raw",
			itemError:   errors.New("negative value"),
		},
	}

	for _, newItem := range scenarios {
		_, err := NewItem(newItem.name, newItem.price, newItem.quantity, newItem.itemType)
		if err != nil && newItem.itemError == nil {
			t.Errorf("For %s got %v  expected was%v", newItem.description, err, newItem.itemError)
		} else if err == nil && newItem.itemError != nil {
			t.Errorf("For %s got %v  expected was%v", newItem.description, err, newItem.itemError)
		}
	}
}

func TestGetEffectivePrice(t *testing.T) {
	scenarios := []scenarioTestEffectivePrice{
		{
			description:   "Manufactured Item",
			item:          Item{Name: "Pen", Price: 12, Quantity: 2, Type: enum.Manufactured},
			expectedPrice: 27.06,
			err:           nil,
		},
		{
			description:   "Raw Item",
			item:          Item{Name: "Book", Price: 26.0, Quantity: 4, Type: enum.Raw},
			expectedPrice: 117.0,
			err:           nil,
		},
		{
			description:   "Imported Item with Rs 5 surplus",
			item:          Item{Name: "Chocolate", Price: 8.0, Quantity: 2, Type: enum.Imported},
			expectedPrice: 22.60,
			err:           nil,
		},
		{
			description:   "Imported Item with Rs 10 surplus",
			item:          Item{Name: "Pencil", Price: 18.0, Quantity: 6, Type: enum.Imported},
			expectedPrice: 128.80,
			err:           nil,
		},
		{
			description:   "Imported Item with  5 %  surplus",
			item:          Item{Name: "Chocolate", Price: 26, Quantity: 8, Type: enum.Imported},
			expectedPrice: 240.240,
			err:           nil,
		},
	}

	for _, testScenarios := range scenarios {
		effectedPriceForItem := testScenarios.item.GetEffectivePrice()
		if effectedPriceForItem != testScenarios.expectedPrice {
			t.Errorf("Error occcuured expected %f got %f for %s", testScenarios.expectedPrice, effectedPriceForItem, testScenarios.description)
		}
	}
}
