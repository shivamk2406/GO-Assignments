package view

import (
	"fmt"
	"os"

	user "github.com/shivamk2406/GO-Assignments/tree/Assignment-2/user"
)

func DriverMenu() {
	var choice int
	TempUser := user.User{}
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
			TempUser.Courses = user.GetCourses()
			TempUser.FullName = user.GetUserName()
			TempUser.Age = (user.GetUserAge())
			TempUser.Address = user.GetAddress()
			TempUser.RollNumber = user.GetRollNumber()
			fmt.Println("Added User")
		case 2:
			fmt.Println("Display User")
			user.DisplayUserDetails()
		case 3:
			fmt.Println("Delete User")
		case 4:
			user.SaveUserDetails(TempUser)
			fmt.Println("Saved User Details")
		case 5:
			os.Exit(1)
			fmt.Println("Exit")
		default:
			fmt.Println("Invalid Input!!! Please enter a valid choice")
		}

	}
}
