package view

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	itemdetails "github.com/shivamk2406/GO-Assignments/item-details"
)

const (
	Accept = "y"
	Deny   = "n"
)

func Initialize() error {
	name, price, quantity, itemType, err := getItem()
	if err != nil {
		return err
	}

	newItem, err := itemdetails.CreateItem(name, price, quantity, itemType)
	if err != nil {
		return err
	}

	fmt.Println(newItem.GetTotalCost())

	moreItems, err := getUserChoice()

	for err != nil {
		moreItems, err = getUserChoice()

	}
	if moreItems == Accept {
		err = Initialize()
		return err
	}
	return nil

}

func getItem() (name string, price float64, quantity int, itemType string, err error) {

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
	return
}

func getUserChoice() (string, error) {
	fmt.Println("Do you want to enter Details of more item:", Accept+"/"+Deny)
	userResponse := Accept
	_, err := fmt.Scanf("%s", &userResponse)
	if err != nil {
		err := errors.Wrap(err, "Scan for user choice failed")
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
	if userResponse != Accept && userResponse != Deny {
		err := fmt.Errorf("invalid Choice")
		log.Println(err)
		return err
	}
	return nil
}
