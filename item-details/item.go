package itemdetails

import (
	"fmt"
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
	enum "github.com/shivamk2406/GO-Assignments/item-details/enum"
)

type Item struct {
	ItemName     string
	ItemPrice    float64
	ItemQuantity int
	ItemType     enum.ItemType
}

func (item Item) CalculateTax(totalCost float64) float64 {
	switch item.ItemType {
	case enum.Raw:
		totalCost += 0.125 * item.ItemPrice * float64(item.ItemQuantity)
	case enum.Manufactured:
		totalCost += 0.125 * item.ItemPrice * float64(item.ItemQuantity)
		totalCost += 0.02 * (item.ItemPrice + 0.125*item.ItemPrice)
	case enum.Imported:
		totalCost = item.applySurcharge(totalCost)
	}
	return totalCost
}

func (item Item) applySurcharge(totalCost float64) float64 {
	totalCost += +0.1 * item.ItemPrice
	if totalCost <= 100 {
		totalCost += 5
	} else if totalCost > 100 && totalCost <= 200 {
		totalCost += 10
	} else {
		totalCost += 0.05 * totalCost
	}
	return totalCost
}

func (item Item) itemInvoice() float64 {
	totalCost := item.ItemPrice * float64(item.ItemQuantity)
	return totalCost

}

func (item Item) GetTotalCost() float64 {

	totalCost := item.itemInvoice()
	fmt.Printf("Cost  for the item without Taxes %f \n", totalCost)

	totalCost = item.CalculateTax(totalCost)
	fmt.Printf("Cost  for the item after Taxes %f \n", totalCost)
	return totalCost
}

func CreateItem(name string, price float64, quantity int, typeItem string) (Item, error) {

	var item Item
	var err error
	item.ItemName = name
	item.ItemPrice = price
	item.ItemQuantity = quantity
	item.ItemType, err = enum.ItemTypeString(typeItem)
	if err != nil {
		return Item{}, err
	}
	err = validateItem(item)
	if err != nil {
		log.Println(err)
		return Item{}, nil
	}
	return item, nil
}

func validateItem(item Item) error {
	return validation.ValidateStruct(&item,
		validation.Field(&item.ItemPrice, validation.By(checkNegativeValue)),
		validation.Field(&item.ItemQuantity, validation.By(checkNegativeValue)))

}

func checkNegativeValue(value interface{}) error {
	err := fmt.Errorf("%v", "negative value")

	switch data := value.(type) {
	case int:
		if data < 0 {
			return err
		}
	case float64:
		if data < 0.0 {
			return err
		}
	}
	return nil

}
