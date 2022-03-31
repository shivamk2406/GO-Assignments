package itemdetails

import "fmt"

//Basic structure for an item
type Item struct {
	ItemName     string
	ItemPrice    float32
	ItemQuantity int
	ItemType     string
}

//Map for item types
/*var itemTypeMap map[string]int = map[string]int{
	"raw":          1,
	"manufactured": 2,
	"imported":     3,
}*/

//Logic for Calculation of Taxes
func (item Item) TaxCalulationLogic(totalCost float32) float32 {
	switch item.ItemType {
	case "raw":
		totalCost = totalCost + 0.125*item.ItemPrice*float32(item.ItemQuantity)
	case "manufactured":
		totalCost = totalCost + 0.125*item.ItemPrice*float32(item.ItemQuantity)
		totalCost = totalCost + 0.02*(item.ItemPrice+0.125*item.ItemPrice)
	case "imported":
		totalCost = totalCost + 0.1*item.ItemPrice
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
func (item Item) GetTotalCost() float32 {

	var totalCost float32
	totalCost = item.ItemPrice * float32(item.ItemQuantity)
	fmt.Printf("Cost  for the item without Taxes %f \n", totalCost)

	totalCost = item.TaxCalulationLogic(totalCost)
	fmt.Printf("Cost  for the item after Taxes %f \n", totalCost)
	return totalCost
}

//New Item Generation
func itemCreated(name string, price float32, quantity int, typeItem string) Item {

	newItem := Item{}

	newItem.ItemName = name
	newItem.ItemPrice = price
	newItem.ItemQuantity = quantity
	newItem.ItemType = typeItem

	return newItem
}
