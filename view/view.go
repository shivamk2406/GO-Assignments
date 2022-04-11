package view

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/aggregate"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/services"
)

func driveMenu() {
	fmt.Println("1.  Add User details.")
	fmt.Println("2.  Display User details.")
	fmt.Println("3.  Delete User details")
	fmt.Println("4.  Save User details.")
	fmt.Println("5.  Exit")

}

func Initialize() error {
	choice := 0
	tempStudents := []aggregate.Student{}

	for choice != 5 {
		driveMenu()
		choice, err := getChoice()
		if err != nil {
			return err
		}
		switch choice {
		case 1:
			tempStudent, err := services.GetStudentDetails()
			if err != nil {
				log.Println(err)
				return err
			}
			tempStudents = append(tempStudents, tempStudent)
			fmt.Println("Student added successfully")
		case 2:
			services.DisplayStudentDetails(tempStudents)
		case 5:
			os.Exit(1)
		}
	}
	return nil
}

func getChoice() (int, error) {
	var choice int
	fmt.Println("Enter Your Choice: ")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		log.Println(err)
		return 0, errors.Errorf("choice scanning failed")
	}

	return choice, nil

}

func validateUserChoice(choice int) error {
	if choice >= 1 && choice <= 5 {
		return errors.Errorf("invalid choice")
	}
	return nil
}

/*func DriverMenu() {
	var choice int
	TempUser := user.User{}
	for choice != 5 {

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
}*/
