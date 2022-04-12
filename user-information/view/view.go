package view

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/aggregate"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/repository"
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
	var rollNo int
	var tempStudents []aggregate.Student

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
			students, err := repository.ReadFromFile()
			if err != nil {
				return err
			}
			fmt.Println("on Temp")
			services.DisplayStudentDetails(tempStudents)
			fmt.Println("on File")
			services.DisplayStudentDetails(students)
		case 3:

			fmt.Println("Enter Roll Number")
			fmt.Scanf("%d", &rollNo)
			err := services.DeleteStudentDetails(uint(rollNo))
			if err != nil {
				return err
			}
		case 4:
			if tempStudents != nil {
				fmt.Println("Saving Details")
				repository.AppendStudentDetails(tempStudents)
			}

			tempStudents = nil

		case 5:
			if tempStudents != nil {
				fmt.Println("There are some unsaved changes!!! Press 1 to save the details")
				choice, err := fmt.Scanf("%d", &choice)
				if err != nil {
					return err
				}
				if choice == 1 {
					repository.AppendStudentDetails(tempStudents)
				}
			}
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
