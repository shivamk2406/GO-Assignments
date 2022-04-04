package firstpage

import (
	"fmt"
	"strings"

	itemdetails "github.com/shivamk2406/GO-Assignments/item-details"
)

func QueryScreen() {
	fmt.Println("Hello There")
	var totalFinalCostOfAllItems float64

	var userResponse string

	isUserInterested := true

	for isUserInterested {
		newItem := itemdetails.Item{}
		newItem.ItemName, newItem.ItemPrice, newItem.ItemQuantity, newItem.ItemType = newItem.

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
