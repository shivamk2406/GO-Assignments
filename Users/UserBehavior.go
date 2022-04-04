package users

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Map for item types.
var itemTypeMap = map[string]int{
	"raw":          1,
	"manufactured": 2,
	"imported":     3,
}

func GetItemInput() (string, float64, int, string) {
	var name string
	var price float64
	var quantity int
	var itemtype string
	fmt.Println("Enter Item Details: ")
	// fmt.Println("Enter First Name of Item")
	name, errorName := getItemName()
	if errorName != nil {
		fmt.Println(errorName)
		os.Exit(1)
	}
	// fmt.Println("Enter Price of Item")
	price, errorPrice := getItemPrice()
	if errorPrice != nil {
		fmt.Println(errorPrice)
		os.Exit(1)
	}
	// fmt.Println("Enter Quantity of Item")
	quantity, errorQuantity := getItemQuantity()
	if errorQuantity != nil {
		fmt.Println(errorQuantity)
		os.Exit(1)
	}
	fmt.Println("Enter Type of Item 1.Raw 2. Manufactured 3.Imported")
	_, errorType := fmt.Scan(&itemtype)
	if errorType != nil {
		fmt.Println(errorType)
		os.Exit(1)
	}

	// All type of lower case and Upper case are converted to lower case in order to map to right category
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

func getItemName() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	var name string
	fmt.Println("Enter First Name of Item")
	scanner.Scan()
	name = scanner.Text()
	isValidName := strings.Contains(name, " ")

	if isValidName {
		err := errors.New("Please Enter First Name only")
		return "", err
	}
	return name, nil
}

func getItemPrice() (float64, error) {
	var price float64
	fmt.Println("Enter Price of an Item")
	fmt.Scan(&price)

	if price < 0 {
		err := errors.New("Negative Price Value not Allowed")
		// err = errors.Wrap(err, "Invalid Value")
		return 0, err
	}

	return price, nil
}

func getItemQuantity() (int, error) {
	var quantityF float64
	fmt.Println("Enter Quantity of item")
	fmt.Scan(&quantityF)

	var quantity int = int(quantityF)

	if quantity < 0 {
		err := errors.New("Negative Quantity Values not Allowed")
		return 0, err
	}
	quantity = int(quantity)
	return quantity, nil
}