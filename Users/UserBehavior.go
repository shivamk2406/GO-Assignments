package users

import (
	"fmt"
	"strings"
	//itemDet "github.com/Blaezy/GO-Assignments/ItemDetails"
)

//Map for item types
var itemTypeMap map[string]int = map[string]int{
	"raw":          1,
	"manufactured": 2,
	"imported":     3,
}

func GetItemInput() (string, float32, int, string) {

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
