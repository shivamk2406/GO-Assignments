package functionalities

import (
	"fmt"
)

func QueryScreen() {
	var choice int
	for choice != 5 {
		fmt.Println("1.  Add User details.")
		fmt.Println("2.  Display User details.")
		fmt.Println("3.  Delete User details")
		fmt.Println("4.  Save User details.")
		fmt.Println("5.  Exit")

		fmt.Println("Enter Your Choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Added User")
		case 2:
			fmt.Println("Display User")
		case 3:
			fmt.Println("Delete User")
		case 4:
			fmt.Println("Save User Details")
		case 5:
			fmt.Println("Exit")
		default:
			fmt.Println("Invalid Input!!! Please enter a valid choice")
		}

	}
}
