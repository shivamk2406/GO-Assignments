package item

import (
	"fmt"
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
	enum "github.com/shivamk2406/GO-Assignments/item/enum"
)

type Item struct {
	Name     string
	Price    float64
	Quantity int
	Type     enum.ItemType
}

type Invoice struct {
	Name           string
	Price          float64
	Quantity       int
	Type           enum.ItemType
	Tax            float64
	EffectivePrice float64
}

const (
	rawItemTaxPerItem                            = 0.125
	baseManufacturedItemTaxPerItem               = 0.125
	addedManufacturedItemTaxPerItem              = 0.02
	rateForAddedManufacturedItemTax              = 112.5
	importedItemTaxPerItem                       = 0.1
	importSurchargeForPriceLessThanHundred       = 5
	importSurchargeForPriceLessThanTwoHundred    = 10
	importSurchargeForPriceGreaterThanTwoHundred = 0.05
)

func (item Item) getItemCostWithoutTax() float64 {
	itemCost := item.Price * float64(item.Quantity)

	return itemCost
}

func (item Item) GetTax() float64 {
	priceWithoutTax := item.getItemCostWithoutTax()
	var tax float64

	switch item.Type {
	case enum.Raw:
		tax = rawItemTaxPerItem * priceWithoutTax
	case enum.Manufactured:
		tax = baseManufacturedItemTaxPerItem * priceWithoutTax
		tax += addedManufacturedItemTaxPerItem * rateForAddedManufacturedItemTax * priceWithoutTax
	case enum.Imported:
		tax = importedItemTaxPerItem * priceWithoutTax
		tax += item.applySurcharge()
	}

	return tax
}

func (item Item) applySurcharge() float64 {
	priceWithoutTax := item.getItemCostWithoutTax()
	priceAfterImportDuty := priceWithoutTax + importedItemTaxPerItem*priceWithoutTax

	if priceAfterImportDuty < 100 {
		return importSurchargeForPriceLessThanHundred
	} else if priceAfterImportDuty >= 100 && priceAfterImportDuty < 200 {
		return importSurchargeForPriceLessThanTwoHundred
	} else {
		return priceAfterImportDuty * importSurchargeForPriceGreaterThanTwoHundred
	}
}

func (item Item) ItemInvoice() Invoice {
	return Invoice{
		Name:           item.Name,
		Price:          item.Price,
		Quantity:       item.Quantity,
		Type:           item.Type,
		Tax:            item.GetTax(),
		EffectivePrice: item.GetFinalPrice(),
	}
}

func (item Item) GetFinalPrice() (price float64) {
	price = item.getItemCostWithoutTax() + item.GetTax()
	return price
}

func NewItem(name string, price float64, quantity int, typeItem string) (Item, error) {
	var item Item
	var err error

	item.Name = name
	item.Price = price
	item.Quantity = quantity
	item.Type, err = enum.ItemTypeString(typeItem)
	if err != nil {
		return Item{}, err
	}

	err = validateItem(item)
	if err != nil {
		log.Println(err)
		return Item{}, err
	}

	return item, nil
}

func validateItem(item Item) error {
	return validation.ValidateStruct(&item,
		validation.Field(&item.Price, validation.By(checkNegativeValue)),
		validation.Field(&item.Quantity, validation.By(checkNegativeValue)))
}

func checkNegativeValue(value interface{}) error {
	switch data := value.(type) {
	case int:
		if data < 0 {
			return fmt.Errorf("negative value")
		}
	case float64:
		if data < 0.0 {
			return fmt.Errorf("negative value")
		}
	}
	return nil
}
