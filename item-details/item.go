package itemdetails

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	ItemName     string
	ItemPrice    float64
	ItemQuantity int
	ItemType     string
}

var itemTypeMap = map[string]int{
	"raw":          1,
	"manufactured": 2,
	"imported":     3,
}

func (item Item) CalculateTax(totalCost float64) float64 {
	switch item.ItemType {
	case "raw":
		totalCost += 0.125 * item.ItemPrice * float64(item.ItemQuantity)
	case "manufactured":
		totalCost += 0.125 * item.ItemPrice * float64(item.ItemQuantity)
		totalCost += 0.02 * (item.ItemPrice + 0.125*item.ItemPrice)
	case "imported":
		totalCost = totalCost + 0.1*item.ItemPrice
		if totalCost <= 100 {
			totalCost = totalCost + 5
		} else if totalCost > 100 && totalCost <= 200 {
			totalCost += 10
		} else {
			totalCost += 0.05 * totalCost
		}
	}
	return totalCost
}

func (item Item) GetTotalCost() float64 {

	var totalCost float64
	totalCost = item.ItemPrice * float64(item.ItemQuantity)
	fmt.Printf("Cost  for the item without Taxes %f \n", totalCost)

	totalCost = item.CalculateTax(totalCost)
	fmt.Printf("Cost  for the item after Taxes %f \n", totalCost)
	return totalCost
}

func GetItemInput() (string, float64, int, string) {
	var name string
	var price float64
	var quantity int
	var itemtype string
	fmt.Println("Enter Item Details: ")

	name, errorName := getItemName()
	if errorName != nil {
		fmt.Println(errorName)
		os.Exit(1)
	}

	price, errorPrice := getItemPrice()
	if errorPrice != nil {
		fmt.Println(errorPrice)
		os.Exit(1)
	}

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
		err := errors.New("please Enter First Name only")
		return "", err
	}
	return name, nil
}

func getItemPrice() (float64, error) {
	var price float64
	fmt.Println("enter Price of an Item")
	fmt.Scan(&price)

	if price < 0 {
		err := errors.New("negative Price Value not Allowed")
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
		err := errors.New("negative Quantity Values not Allowed")
		return 0, err
	}
	quantity = int(quantity)
	return quantity, nil
}

// New Item Generation
/*func itemCreated(name string, price float64, quantity int, typeItem string) Item {

	newItem := Item{}

	newItem.ItemName = name
	newItem.ItemPrice = price
	newItem.ItemQuantity = quantity
	newItem.ItemType = typeItem

	return newItem
}*/
