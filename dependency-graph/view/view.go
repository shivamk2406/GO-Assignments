package view

import (
	"fmt"
	"log"

	node "github.com/shivamk2406/dependency-graph/domain/graph"
	"golang.org/x/exp/maps"
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

func populateGraph(familyTree node.FamilyTree) {
	familyTree.AddNode(1, "A")
	familyTree.AddNode(2, "B")
	familyTree.AddNode(3, "C")
	familyTree.AddNode(4, "D")
	familyTree.AddNode(5, "E")
	familyTree.AddNode(6, "F")
	familyTree.AddNode(7, "G")

	familyTree.AddEdge(1, 3)
	familyTree.AddEdge(1, 4)
	familyTree.AddEdge(1, 5)
	familyTree.AddEdge(2, 3)
	familyTree.AddEdge(2, 4)
	familyTree.AddEdge(2, 5)
	familyTree.AddEdge(3, 6)
	familyTree.AddEdge(3, 7)
	familyTree.AddEdge(4, 6)
	familyTree.AddEdge(4, 7)
	familyTree.AddEdge(5, 6)
	familyTree.AddEdge(5, 7)

}

func Initialize() error {
	var choice int
	var err error
	familyTree := node.NewFamilyTree()
	populateGraph(familyTree)
	for choice != int(ExitChoice) {
		showMenu()
		choice, err = getUserChoice()
		if err != nil {
			return err
		}

		switch choice {
		case int(ImmediateParentsChoice):
			err := getImmediateParents(familyTree)
			if err != nil {
				return err
			}
		case int(ImmediateChildrenChoice):
			err := getImmediateChildren(familyTree)
			if err != nil {
				return err
			}
		case int(AncestorsChoice):
			err := getAncestors(familyTree)
			if err != nil {
				return err
			}
		case int(DescendentsChoice):
			err := getDescendents(familyTree)
			if err != nil {
				return err
			}
		case int(DeleteDependencyChoice):
			err := deleteDependency(familyTree)
			if err != nil {
				return err
			}
		case int(DeleteNodeChoice):
			err := deleteNode(familyTree)
			if err != nil {
				return err
			}
		case int(AddDependencyChoice):
			err := addDependency(familyTree)
			if err != nil {
				return err
			}
		case int(AddNodeChoice):
			err := addNewNode(familyTree)
			if err != nil {
				return err
			}
		case int(ExitChoice):
			break
		default:
			fmt.Println("Invalid choice please try again")
		}
	}
	fmt.Println("Exiting")
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

func getAncestors(familyTree node.FamilyTree) error {
	var id int

	fmt.Println("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return err
	}

	ancestors, err := familyTree.GetAncestors(id)
	if err != nil {
		return err
	}

	keys := maps.Keys(ancestors)
	fmt.Println("The Ancestors are:")
	for i := range keys {
		fmt.Printf("%d  ", keys[i])
	}
	fmt.Println()

	return nil
}

func getDescendents(familyTree node.FamilyTree) error {
	var id int

	fmt.Println("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return err
	}

	descendants, err := familyTree.GetDescendents(id)
	if err != nil {
		return err
	}

	keys := maps.Keys(descendants)
	fmt.Println("The Descendants are:")
	for i := range keys {
		fmt.Printf("%d  ", keys[i])
	}
	fmt.Println()

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
