package item

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/config"
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

func (item Item) getBasePrice() float64 {
	itemCost := item.Price * float64(item.Quantity)

	return itemCost
}

func (item Item) GetTax() float64 {
	switch item.Type {
	case enum.Raw:
		return config.RawItemTaxRate * item.getBasePrice()
	case enum.Manufactured:
		baseTax := config.BaseManufacturedItemTaxRate * item.getBasePrice()
		return baseTax + config.AddedManufacturedItemTaxRate*baseTax
	case enum.Imported:
		return config.ImportedItemTaxRate*item.getBasePrice() + item.applySurcharge()
	}
	return 0
}

func (item Item) applySurcharge() float64 {
	importedPrice := item.getBasePrice() + config.ImportedItemTaxRate*item.getBasePrice()
	if importedPrice < config.SurchargeLimit1 {
		return config.Limit1Surcharge
	} else if importedPrice >= config.SurchargeLimit1 && importedPrice < config.SurchargeLimit2 {
		return config.Limit2Surcharge
	} else {
		return importedPrice * config.Limit3SurchargeRate
	}
}

func (item Item) ItemInvoice() Invoice {
	return Invoice{
		Name:           item.Name,
		Price:          item.Price,
		Quantity:       item.Quantity,
		Type:           item.Type,
		Tax:            item.GetTax(),
		EffectivePrice: item.GetEffectivePrice(),
	}
}

func (item Item) GetEffectivePrice() (price float64) {
	price = item.getBasePrice() + item.GetTax()
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
			return errors.Errorf("negative value")
		}
	case float64:
		if data < 0.0 {
			return errors.Errorf("negative value")
		}
	}
	return nil
}
