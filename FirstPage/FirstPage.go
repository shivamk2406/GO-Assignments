package firstpage

import (
	"fmt"
	"strings"

	itemdetails "github.com/Blaezy/GO-Assignments/ItemDetails"
	users "github.com/Blaezy/GO-Assignments/Users"
)

func QueryScreen() {
	fmt.Println("Hello There")

	fmt.Println("Hello There")
	var totalFinalCostOfAllItems float64

	//User Response variable is used to store response corresponding to next item addition
	var userResponse string
	//boolean variable to check whether user is interested in adding items or not
	isUserInterested := true

	for isUserInterested {
		newItem := itemdetails.Item{}
		newItem.ItemName, newItem.ItemPrice, newItem.ItemQuantity, newItem.ItemType = users.GetItemInput()

		//newItem := Item{itemName: name, itemPrice: price, itemQuantity: quantity, itemType: itemtype}
		//fmt.Println(newItem.getTotalCost())
		totalFinalCostOfAllItems = totalFinalCostOfAllItems + newItem.GetTotalCost()
		fmt.Println("Do you want to enter details of any other item (y/n):")
		fmt.Scan(&userResponse)
		userResponse = strings.ToLower(userResponse)

		if userResponse == "n" {
			isUserInterested = false
		} else {
			//fmt.Println("Invalid Response!!!! Try Again")
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
