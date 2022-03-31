package firstpage

import (
	"fmt"

	itemdetails "github.com/Blaezy/GO-Assignments/ItemDetails"
	users "github.com/Blaezy/GO-Assignments/Users"
)

func QueryScreen() {
	fmt.Println("Hello There")
	newItem := itemdetails.Item{}
	newItem.ItemName, newItem.ItemPrice, newItem.ItemQuantity, newItem.ItemType = users.GetItemInput()

}
