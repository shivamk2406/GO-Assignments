package main

import (
	"fmt"
	"strings"
)

//Basic structure for an item
type Item struct {
	itemName     string
	itemPrice    float32
	itemQuantity int
	itemType     string
}

//Map for item types
var itemTypeMap map[string]int = map[string]int{
	"raw":          1,
	"manufactured": 2,
	"imported":     3,
}

//Logic for Calculation of Taxes
func (item Item) taxCalulationLogic(totalCost float32) float32 {
	switch item.itemType {
	case "raw":
		totalCost = totalCost + 0.125*item.itemPrice*float32(item.itemQuantity)
	case "manufactured":
		totalCost = totalCost + 0.125*item.itemPrice*float32(item.itemQuantity)
		totalCost = totalCost + 0.02*(item.itemPrice+0.125*item.itemPrice)
	case "imported":
		totalCost = totalCost + 0.1*item.itemPrice
		if totalCost <= 100 {
			totalCost = totalCost + 5
		} else if totalCost > 100 && totalCost <= 200 {
			totalCost = totalCost + 10
		} else {
			totalCost = totalCost + 0.05*totalCost
		}
	}
	return totalCost
}

//Total Cost before and after Tax Calculation
func (item Item) getTotalCost() float32 {

	var totalCost float32
	totalCost = item.itemPrice * float32(item.itemQuantity)
	fmt.Printf("Cost  for the item without Taxes %f \n", totalCost)

	totalCost = item.taxCalulationLogic(totalCost)
	fmt.Printf("Cost  for the item after Taxes %f \n", totalCost)
	return totalCost
}

/*func main() {
	fmt.Println("Hello There")
	var totalFinalCostOfAllItems float32

	//User Response variable is used to store response corresponding to next item addition
	var userResponse string
	//boolean variable to check whether user is interested in adding items or not
	isUserInterested := true

	for isUserInterested {
		newItem := Item{}
		newItem.itemName, newItem.itemPrice, newItem.itemQuantity, newItem.itemType = getItemInput()
		//newItem := Item{itemName: name, itemPrice: price, itemQuantity: quantity, itemType: itemtype}
		//fmt.Println(newItem.getTotalCost())
		totalFinalCostOfAllItems = totalFinalCostOfAllItems + newItem.getTotalCost()
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
}*/

//function to get all item related information from the user
func getItemInput() (string, float32, int, string) {

	var name string
	var price float32
	var quantity int
	var itemtype string
	fmt.Println("Enter Item Details: ")
	fmt.Println("Enter First Name of Item")
	_, errorName := fmt.Scan(&name)
	if errorName != nil {
		fmt.Println(errorName)
	}
	fmt.Println("Enter Price of Item")
	_, errorPrice := fmt.Scan(&price)
	if errorPrice != nil {
		fmt.Println(errorPrice)
	}
	fmt.Println("Enter Quantity of Item")
	_, errorQuantity := fmt.Scanln(&quantity)
	if errorQuantity != nil {
		fmt.Println(errorQuantity)
	}
	fmt.Println("Enter Type of Item 1.Raw 2. Manufactured 3.Imported")
	_, errorType := fmt.Scan(&itemtype)
	if errorType != nil {
		fmt.Println(errorType)
	}

	//All type of lower case and Upper case are converted to lower case in order to map to right category
	itemtype = strings.ToLower(itemtype)
	_, isValid := itemTypeMap[itemtype]

	for !isValid {
		fmt.Println("Invalid Type!!!")
		fmt.Println("Enter Type of Item 1.Raw 2. Manufactured 3.Imported")
		fmt.Scan(&itemtype)
		_, isValid := itemTypeMap[itemtype]
		if isValid {
			break
		}
	}
	return name, price, quantity, itemtype
}
