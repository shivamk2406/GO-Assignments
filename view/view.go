package view

import (
	"fmt"
	"log"

	itemdetails "github.com/shivamk2406/GO-Assignments/item"
	"github.com/shivamk2406/GO-Assignments/item/enum"
)

func Initialize() error {
	item, err := getItem()
	if err != nil {
		return err
	}

	itemInvoice := item.ItemInvoice()
	fmt.Println("Item Details:")
	fmt.Printf("Name: %s\nPrice: %f \nQuantity: %d \nType: %s \nTax: %f \nEffective Price: %f\n", itemInvoice.Name, itemInvoice.Price, item.Quantity, itemInvoice.Type, itemInvoice.Tax, itemInvoice.EffectivePrice)

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

func getItem() (itemdetails.Item, error) {
	var name string
	var price float64
	var quantity int
	var itemType string

	fmt.Println("Enter Item Name")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		log.Println(err)
		return itemdetails.Item{}, fmt.Errorf("item name scanning failed")
	}

	fmt.Println("Enter Item Price")
	_, err = fmt.Scanf("%f", &price)
	if err != nil {
		log.Println(err)
		return itemdetails.Item{}, fmt.Errorf("item price scanning failed")
	}

	fmt.Println("Enter Item Quantity")
	_, err = fmt.Scanf("%d", &quantity)
	if err != nil {
		log.Println(err)
		return itemdetails.Item{}, fmt.Errorf("item quantity scanning failed")
	}

	fmt.Println("Enter Item Type")
	_, err = fmt.Scanf("%s", &itemType)
	if err != nil {
		log.Println(err)
		return itemdetails.Item{}, fmt.Errorf("item quantity scanning failed")
	}

	item, err := itemdetails.NewItem(name, price, quantity, itemType)
	if err != nil {
		return itemdetails.Item{}, err
	}

	return item, nil
}

func getUserChoice() (string, error) {
	fmt.Println("Do you want to enter Details of more item:", enum.Accept+"/"+enum.Deny)
	userResponse := enum.Accept
	_, err := fmt.Scanf("%s", &userResponse)
	if err != nil {
		log.Println(err)
		return userResponse, fmt.Errorf("scan for user chhoice failed")
	}

	if err := validateUserResponse(userResponse); err != nil {

		return userResponse, fmt.Errorf("invalid user response")
	}

	return userResponse, nil
}

func validateUserResponse(userResponse string) error {
	if userResponse != enum.Accept && userResponse != enum.Deny {
		return fmt.Errorf("invalid Choice")
	}
	return nil
}
