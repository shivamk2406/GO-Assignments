package item

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"github.com/shivamk2406/item-inventory/domain/item/enum"
)

const (
	Accept                          = "y"
	Deny                            = "n"
	RawItemTaxRate                  = 0.125
	BaseManufacturedItemTaxRate     = 0.125
	AddedManufacturedItemTaxRate    = 0.02
	RateForAddedManufacturedItemTax = 112.5
	ImportedItemTaxRate             = 0.1
	Limit1Surcharge                 = 5
	Limit2Surcharge                 = 10
	Limit3SurchargeRate             = 0.05
	SurchargeLimit1                 = 100
	SurchargeLimit2                 = 200
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
		return RawItemTaxRate * item.getBasePrice()
	case enum.Manufactured:
		baseTax := BaseManufacturedItemTaxRate * item.getBasePrice()
		return baseTax + AddedManufacturedItemTaxRate*baseTax
	case enum.Imported:
		return ImportedItemTaxRate*item.getBasePrice() + item.applySurcharge()
	}
	return 0
}

func (item Item) applySurcharge() float64 {
	importedPrice := item.getBasePrice() + ImportedItemTaxRate*item.getBasePrice()
	if importedPrice < SurchargeLimit1 {
		return Limit1Surcharge
	} else if importedPrice >= SurchargeLimit1 && importedPrice < SurchargeLimit2 {
		return Limit2Surcharge
	} else {
		return importedPrice * Limit3SurchargeRate
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
