package view

import (
	"fmt"
	"log"
	"os"
)

func showMenu() {
	fmt.Println("****Operations Available****")
	fmt.Println("1.Get the immediate parents of a node")
	fmt.Println("2.Get the immediate children of a node")
	fmt.Println("3.Get the ancestors of a node")
	fmt.Println("4.Get the descendants of a node")
	fmt.Println("5.Delete dependency from a tree")
	fmt.Println("6.Delete a node from a tree")
	fmt.Println("7.Add a new dependency to a tree")
	fmt.Println("8.Add a new node to tree")
	fmt.Println("9. exit")
}

func Initialize() error {
	var choice int
	var err error

	for choice != 9 {
		choice, err = getUserChoice()
		if err != nil {
			return err
		}
		showMenu()
		switch choice {
		case 1:
			fmt.Println("Pressed 1")
		case 9:
			os.Exit(1)
		}
	}
	return nil
}

func getUserChoice() (int, error) {
	var choice int
	fmt.Println("Enter your choice")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return choice, err
}
