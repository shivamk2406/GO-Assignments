package view

import (
	"fmt"
	"log"
	"os"

	"github.com/shivamk2406/dependency-graph/domain/node"
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
	familyTree := node.NewFamilyTree()

	for choice != 9 {
		showMenu()
		choice, err = getUserChoice()
		if err != nil {
			return err
		}

		switch choice {
		case 1:
			err := getImmediateParents(familyTree)
			if err != nil {
				return err
			}
		case 2:
			err := getImmediateChildren(familyTree)
			if err != nil {
				return err
			}
		case 5:
			err := deleteDependency(familyTree)
			if err != nil {
				return err
			}
		case 6:
			err := deleteNode(familyTree)
			if err != nil {
				return err
			}
		case 7:
			err := addDependency(familyTree)
			if err != nil {
				return err
			}
		case 8:
			err := addNewNode(familyTree)
			if err != nil {
				return err
			}
		case 9:
			os.Exit(1)
		}
	}
	return nil
}

func getImmediateParents(familyTree node.FamilyTree) error {
	var id int
	fmt.Printf("Enter Node Id:")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return err
	}

	parents, err := familyTree.GetParents(id)
	if err != nil {
		return err
	}

	for id := range parents {
		fmt.Println(id)
	}
	return nil
}

func getImmediateChildren(familyTree node.FamilyTree) error {
	var id int
	fmt.Printf("Enter Node Id:")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return err
	}

	children, err := familyTree.GetChildren(id)
	if err != nil {
		return err
	}

	for id := range children {
		fmt.Println(id)
	}
	return nil
}

func deleteDependency(familyTree node.FamilyTree) error {
	var id1 int
	var id2 int
	fmt.Printf("Enter id-1: ")
	_, err := fmt.Scanf("%d", &id1)
	if err != nil {
		return err
	}

	fmt.Printf("Enter id-2: ")
	_, err = fmt.Scanf("%d", &id2)
	if err != nil {
		return err
	}

	err = familyTree.DeleteEdge(id1, id2)
	if err != nil {
		return err
	}

	fmt.Println("dependency deletion successful")
	return nil
}

func deleteNode(familyTree node.FamilyTree) error {
	var id int
	fmt.Printf("Enter Node Id:")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return err
	}
	err = familyTree.DeleteNode(id)
	if err != nil {
		return err
	}

	fmt.Println("node deletion successful")
	return nil
}

func addDependency(familyTree node.FamilyTree) error {
	var id1 int
	var id2 int
	fmt.Printf("Enter id-1: ")
	_, err := fmt.Scanf("%d", &id1)
	if err != nil {
		return err
	}

	fmt.Printf("Enter id-2: ")
	_, err = fmt.Scanf("%d", &id2)
	if err != nil {
		return err
	}

	err = familyTree.AddEdge(id1, id2)
	if err != nil {
		return err
	}

	fmt.Println("dependency addition successful")
	return nil
}

func addNewNode(familyTree node.FamilyTree) error {
	var id int
	var name string

	fmt.Println("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return err
	}

	fmt.Printf("Enter name: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		return err
	}

	err = familyTree.AddNode(id, name)
	if err != nil {
		return err
	}

	fmt.Println("node addition successful")
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
