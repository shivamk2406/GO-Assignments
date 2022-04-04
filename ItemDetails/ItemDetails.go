package ItemDetails

import "fmt"

// Basic structure for an item
type Item struct {
	ItemName     string
	ItemPrice    float64
	ItemQuantity int
	ItemType     string
}

// Map for item types
/*var itemTypeMap map[string]int = map[string]int{
	"raw":          1,
	"manufactured": 2,
	"imported":     3,
}*/

// Logic for Calculation of Taxes
func (item Item) TaxCalulationLogic(totalCost float64) float64 {
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

// Total Cost before and after Tax Calculation
func (item Item) GetTotalCost() float64 {

	var totalCost float64
	totalCost = item.ItemPrice * float64(item.ItemQuantity)
	fmt.Printf("Cost  for the item without Taxes %f \n", totalCost)

	totalCost = item.TaxCalulationLogic(totalCost)
	fmt.Printf("Cost  for the item after Taxes %f \n", totalCost)
	return totalCost
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
