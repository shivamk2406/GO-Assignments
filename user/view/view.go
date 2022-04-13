package view

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/domain/students"
)

const (
	Agree = "y"
	Deny  = "n"
)

func showMenu() {
	fmt.Println("1.  Add User details.")
	fmt.Println("2.  Display User details.")
	fmt.Println("3.  Delete User details")
	fmt.Println("4.  Save User details.")
	fmt.Println("5.  Exit")

}

func addUserDetails() (students.Student, error) {
	tempStudent, err := students.GetStudentDetails()
	if err != nil {
		log.Println(err)
		return students.Student{}, err
	}
	err = students.ValidateDuplicates(tempStudent)
	if err != nil {
		return students.Student{}, err
	}
	return tempStudent, nil
}

func confirmSave(tempStudents []students.Student) error {
	var choice string
	fmt.Println("There are some unsaved changes!!! Do you want save those changes", Agree+"/"+Deny)
	_, err := fmt.Scanf("%s", &choice)
	if err != nil {
		return err
	}
	if err := validateUserResponse(choice); err != nil {
		return err
	}
	if choice == Agree {
		err := students.AppendStudentDetails(tempStudents)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Exiting")
	}
	return nil
}

func validateUserResponse(userResponse string) error {
	if userResponse != Agree && userResponse != Deny {
		return errors.Errorf("invalid Choice")
	}
	return nil
}

func Initialize() error {
	choice := 0
	var tempStudents []students.Student

	for choice != 5 {
		showMenu()
		choice, err := getChoice()
		if err != nil {
			return err
		}
		switch choice {
		case 1:
			tempStudent, err := addUserDetails()
			if err != nil {
				fmt.Println(err)
			}
			tempStudents = append(tempStudents, tempStudent)
			fmt.Println("Student added successfully")
		case 2:
			existingStudents, err := students.ReadFromFile()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("on Temp")
			fmt.Println("----------------------------------------------------------------------------------")
			students.DisplayStudentDetails(tempStudents)
			fmt.Println("on File")
			fmt.Println("----------------------------------------------------------------------------------")
			students.DisplayStudentDetails(existingStudents)
		case 3:
			err := students.DeleteStudentDetails()
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			if tempStudents != nil {
				fmt.Println("Saving Details")
				err := students.AppendStudentDetails(tempStudents)
				if err != nil {
					fmt.Println(err)
				}
			}
			tempStudents = nil

		case 5:
			if tempStudents != nil {
				err := confirmSave(tempStudents)
				if err != nil {
					fmt.Println(err)
				}
			}
			os.Exit(1)
		default:
			fmt.Println("invalid choice")
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
