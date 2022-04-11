package view

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	itemdetails "github.com/shivamk2406/GO-Assignments/item"
	"github.com/shivamk2406/GO-Assignments/item/enum"
)

func Initialize() error {
	item, err := getItem()
	if err != nil {
		return err
	}

	itemInvoice := item.ItemInvoice()
	fmt.Printf("Item Details:\nName: %s\nPrice: %f \nQuantity: %d \nType: %s \nTax: %f \nEffective Price: %f\n", itemInvoice.Name, itemInvoice.Price, item.Quantity, itemInvoice.Type, itemInvoice.Tax, itemInvoice.EffectivePrice)

	moreItems, err := getUserChoice()

	for err != nil {
		moreItems, err = getUserChoice()
	}

	if moreItems == enum.Accept {
		err = Initialize()
		return err
	}
	return nil
}

func getItem() (item itemdetails.Item, err error) {
	var name string
	var price float64
	var quantity int
	var itemType string

	fmt.Println("Enter Item Name")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		err = errors.Wrap(err, "item name scanning failed")
		log.Println(err)
		return
	}

	fmt.Println("Enter Item Price")
	_, err = fmt.Scanf("%f", &price)
	if err != nil {
		err = errors.Wrap(err, "item name scanning failed")
		log.Println(err)
		return
	}

	fmt.Println("Enter Item Quantity")
	_, err = fmt.Scanf("%d", &quantity)
	if err != nil {
		err = errors.Wrap(err, "item quantity scanning failed")
		log.Println(err)
		return
	}

	fmt.Println("Enter Item Type")
	_, err = fmt.Scanf("%s", &itemType)
	if err != nil {
		err = errors.Wrap(err, "item type scanning failed")
		log.Println(err)
		return
	}

	item, err = itemdetails.NewItem(name, price, quantity, itemType)
	if err != nil {
		return
	}

	return
}

func getUserChoice() (string, error) {
	fmt.Println("Do you want to enter Details of more item:", enum.Accept+"/"+enum.Deny)
	userResponse := enum.Accept
	_, err := fmt.Scanf("%s", &userResponse)
	if err != nil {
		err = errors.Wrap(err, "Scan for user choice failed")
		log.Println(err)
		return userResponse, err
	}

	if err := validateUserResponse(userResponse); err != nil {
		err = errors.Wrap(err, "invalid user response")
		return userResponse, err
	}

	return userResponse, nil
}

func validateUserResponse(userResponse string) error {
	if userResponse != enum.Accept && userResponse != enum.Deny {
		return fmt.Errorf("invalid Choice")
	}
	return nil
}
