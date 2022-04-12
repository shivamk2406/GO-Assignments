package view

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/aggregate"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/repository"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/services"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/vutil"
)

func showMenu() {
	fmt.Println("1.  Add User details.")
	fmt.Println("2.  Display User details.")
	fmt.Println("3.  Delete User details")
	fmt.Println("4.  Save User details.")
	fmt.Println("5.  Exit")

}

func addUserDetails() (aggregate.Student, error) {
	tempStudent, err := services.GetStudentDetails()
	if err != nil {
		log.Println(err)
		return aggregate.Student{}, err
	}
	err = vutil.CheckDuplicates(tempStudent)
	if err != nil {
		return aggregate.Student{}, err
	}
	return tempStudent, nil
}

func safeExit(tempStudents []aggregate.Student) error {
	var choice int
	fmt.Println("There are some unsaved changes!!! Press 1 to save the details any key to exit")
	choice, err := fmt.Scanf("%d", &choice)
	if err != nil {
		return err
	}

	if choice == 1 {
		err := repository.AppendStudentDetails(tempStudents)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Exiting")
	}
	return nil
}

func Initialize() error {
	choice := 0
	var tempStudents []aggregate.Student

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
				return err
			}
			tempStudents = append(tempStudents, tempStudent)
			fmt.Println("Student added successfully")
		case 2:
			students, err := repository.ReadFromFile()
			if err != nil {
				return err
			}
			fmt.Println("on Temp")
			fmt.Println("----------------------------------------------------------------------------------")
			services.DisplayStudentDetails(tempStudents)
			fmt.Println("on File")
			fmt.Println("----------------------------------------------------------------------------------")
			services.DisplayStudentDetails(students)
		case 3:
			err := services.DeleteStudentDetails()
			if err != nil {
				return err
			}
		case 4:
			if tempStudents != nil {
				fmt.Println("Saving Details")
				err := repository.AppendStudentDetails(tempStudents)
				if err != nil {
					return err
				}
			}
			tempStudents = nil

		case 5:
			if tempStudents != nil {
				err := safeExit(tempStudents)
				if err != nil {
					return err
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
