package firstpage

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	itemdetails "github.com/shivamk2406/GO-Assignments/item-details"
)

func QueryScreen() {
	fmt.Println("Hello There")
	var totalFinalCostOfAllItems float64

	var userResponse string

	isUserInterested := true

	for isUserInterested {
		newItem := itemdetails.Item{}
		newItem.ItemName, newItem.ItemPrice, newItem.ItemQuantity, newItem.ItemType = itemdetails.GetItemInput()

		totalFinalCostOfAllItems += newItem.GetTotalCost()
		fmt.Println("Do you want to enter details of any other item (y/n):")
		fmt.Scan(&userResponse)
		userResponse = strings.ToLower(userResponse)

		if userResponse == "n" {
			isUserInterested = false
		} else {
			// fmt.Println("Invalid Response!!!! Try Again")
			for !(userResponse == "y") && !(userResponse == "n") {
				fmt.Println("Invalid Response Try Again!!!!!")
				fmt.Println("Do you want to enter details of any other item (y/n):")
				fmt.Scan(&userResponse)
				userResponse = strings.ToLower(userResponse)
				if userResponse == "n" {
					isUserInterested = false
				}
			}
		}
	}
	fmt.Println("Total Cost of all Items Including Taxes are: ", totalFinalCostOfAllItems)

}

func Initialize() err {
	name, price, quantity, itemType, err := getItem()
	if err!=nil{
		return err
	}

	newItem,err:=itemdetails.

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
