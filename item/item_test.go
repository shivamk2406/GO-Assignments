package item

import (
	"errors"
	"testing"
)

type scenarioTestNewItem struct {
	description string
	name        string
	price       float64
	quantity    int
	itemType    string
	itemError   error
}

func TestNewItem(t *testing.T) {
	var scenarios = []scenarioTestNewItem{
		{description: "All information is provided and Price is negative",
			name:      "Pen",
			price:     -87,
			quantity:  4,
			itemType:  "raw",
			itemError: errors.New("negative value."),
		},
		{description: "All correct information is provided",
			name:      "Book",
			price:     56,
			quantity:  4,
			itemType:  "manufactured",
			itemError: nil,
		},
		{description: "All information is provided and quantity is negative",
			name:      "Pen",
			price:     87,
			quantity:  -4,
			itemType:  "raw",
			itemError: errors.New("negative value."),
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
