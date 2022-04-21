package view

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/domain/students"
)

const (
	Agree       = "y"
	Deny        = "n"
	AddUser     = 1
	DisplayUser = 2
	DeleteUser  = 3
	SaveUser    = 4
	ExitMenu    = 5
)

func Initialize() error {
	studentRepo := students.NewRepo()
	err := studentRepo.Load()
	if err != nil {
		return err
	}
	moreInput := true

	for moreInput {
		displayMenu()
		choice, err := getChoice()
		if err != nil {
			return err
		}

		switch choice {
		case AddUser:
			err = addUser(studentRepo)
			if err != nil {
				log.Println(err)
			}
		case DisplayUser:
			err = displayUser(studentRepo)
			if err != nil {
				log.Println(err)
			}
		case DeleteUser:
			err = deleteUser(studentRepo)
			if err != nil {
				log.Println(err)
			}
		case SaveUser:
			err = saveUser(studentRepo)
			if err != nil {
				log.Println(err)
			}
		case ExitMenu:
			moreInput = false
			err = confirmSave(studentRepo)
			if err != nil {
				log.Println(err)
			}
			os.Exit(1)
		default:
			fmt.Println("invalid choice")
		}
	}
	return nil
}

func displayMenu() {
	fmt.Println("1.  Add User details.")
	fmt.Println("2.  Display User details.")
	fmt.Println("3.  Delete User details")
	fmt.Println("4.  Save User details.")
	fmt.Println("5.  Exit")
}

func addUser(userRepo students.Repository) error {
	tempStudent, err := students.GetStudent()
	if err != nil {
		log.Println(err)
		return err
	}

	err1 := userRepo.Add(tempStudent)
	if err1 != nil {
		return err1
	}

	fmt.Println("user added successfully")
	return nil
}

func displayUser(userRepo students.Repository) error {
	return userRepo.Display()
}

func deleteUser(userRepo students.Repository) error {
	var rollNumber int
	fmt.Println("Enter Roll Number:")
	fmt.Scanf("%d", &rollNumber)
	err := userRepo.Delete(rollNumber)
	if err != nil {
		return err
	}

	return nil
}

func saveUser(userRepo students.Repository) error {
	err := userRepo.Save()
	if err != nil {
		return err
	}

	fmt.Println("Saved Successfully!!")
	return nil
}

func confirmSave(userRepo students.Repository) error {
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
		err := userRepo.Save()
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
