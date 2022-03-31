package itemdetails

import "fmt"

//Basic structure for an item
type Item struct {
	itemName     string
	itemPrice    float32
	itemQuantity int
	itemType     string
}

//Map for item types
/*var itemTypeMap map[string]int = map[string]int{
	"raw":          1,
	"manufactured": 2,
	"imported":     3,
}*/

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

//New Item Generation
func itemCreated(name string, price float32, quantity int, typeItem string) Item {

	newItem := Item{}

	newItem.itemName = name
	newItem.itemPrice = price
	newItem.itemQuantity = quantity
	newItem.itemType = typeItem

	return newItem
}
