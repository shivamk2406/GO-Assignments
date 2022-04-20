package node

import (
	"fmt"
)

type FamilyTree interface {
	AddNode(int, string) error
	AddEdge(int, int) error
	DeleteNode(int) error
	DeleteEdge(int, int) error
	GetParents(int) (map[int]*Node, error)
	GetChildren(int) (map[int]*Node, error)
	GetAncestors(int) (map[int]*Node, error)
	GetDescendents(int) (map[int]*Node, error)
}

type familyTree struct {
	nodes map[int]*Node
}

func NewFamilyTree() *familyTree {
	return &familyTree{}
}

type Node struct {
	Id       int
	Name     string
	Parents  map[int]*Node
	Children map[int]*Node
}

func (f *familyTree) AddNode(id int, name string) error {
	if f.nodes == nil {
		f.nodes = make(map[int]*Node)
	}

	_, exists := f.nodes[id]
	if exists {
		return fmt.Errorf("node already exists")
	}

	f.nodes[id] = &Node{Id: id, Name: name}
	return nil
}

func (f *familyTree) AddEdge(id1 int, id2 int) error {
	if f.nodes == nil {
		return fmt.Errorf("nodes do not exist")
	}

	_, exists := f.nodes[id1]
	if !exists {
		return fmt.Errorf("node" + string(rune(id1)) + "do not exists")
	}

	_, exists = f.nodes[id2]
	if !exists {
		return fmt.Errorf("node" + string(rune(id2)) + "do not exists")
	}

	if f.nodes[id1].Children == nil {
		f.nodes[id1].Children = make(map[int]*Node)
	}

	//make id2 as child of id1
	f.nodes[id1].Children[id2] = f.nodes[id2]

	if f.nodes[id2].Parents == nil {
		f.nodes[id2].Parents = make(map[int]*Node)
	}
	//make id1 as parent of id2
	f.nodes[id2].Parents[id1] = f.nodes[id1]

	return nil
}

func (f *familyTree) DeleteNode(id int) error {
	if f.nodes == nil {
		return fmt.Errorf("no nodes in the graph")
	}

	_, exists := f.nodes[id]
	if !exists {
		return fmt.Errorf("no such node exists")
	}

	child := f.nodes[id].Children
	for _, node := range child {
		parent := node.Parents
		delete(parent, id)
	}

	parents := f.nodes[id].Parents
	for _, node := range parents {
		children := node.Children
		delete(children, id)
	}

	return nil
}

func (f *familyTree) DeleteEdge(id1 int, id2 int) error {
	if f.nodes == nil {
		return fmt.Errorf("nodes do not exist")
	}

	_, exists := f.nodes[id1]
	if !exists {
		return fmt.Errorf("node" + string(rune(id1)) + "do not exists")
	}

	_, exists = f.nodes[id2]
	if !exists {
		return fmt.Errorf("node" + string(rune(id2)) + "do not exists")
	}

	children := f.nodes[id1].Children
	if children == nil {
		return fmt.Errorf("no children exists for the node")
	}

	_, exists = children[id2]
	if !exists {
		return fmt.Errorf("no relationship found between the two nodes")
	}

	parents := f.nodes[id2].Parents

	delete(parents, id1)
	delete(children, id2)

	return nil
}

func (f *familyTree) GetChildren(id int) (map[int]*Node, error) {
	if f.nodes == nil {
		return nil, fmt.Errorf("no nodes exists in graph ")
	}

	_, exists := f.nodes[id]
	if !exists {
		return nil, fmt.Errorf("no such nodes exists")
	}

	children := make(map[int]*Node)

	for id, node := range f.nodes[id].Children {
		children[id] = node
	}

	return children, nil
}

func (f *familyTree) GetParents(id int) (map[int]*Node, error) {
	if f.nodes == nil {
		return nil, fmt.Errorf("no nodes exists in graph ")
	}

	_, exists := f.nodes[id]
	if !exists {
		return nil, fmt.Errorf("no such nodes exists")
	}
	parents := make(map[int]*Node)

	for id, node := range f.nodes[id].Parents {
		parents[id] = node
	}

	return parents, nil
}

func (f *familyTree) GetAncestors(id int) (map[int]*Node, error) {
	if f.nodes == nil {
		return nil, fmt.Errorf("no nodes exists in graph ")
	}

	_, exists := f.nodes[id]
	if !exists {
		return nil, fmt.Errorf("no such nodes exists")
	}

	ancestors := make(map[int]*Node)
	//visited := make(map[int]bool)

	getAnc(f.nodes[id].Parents, ancestors)
	return ancestors, nil
}

func getAnc(parents map[int]*Node, ancestors map[int]*Node) {
	for _, val := range parents {
		_, exists := ancestors[val.Id]
		if !exists {
			ancestors[val.Id] = val
		}

	}
	for _, val := range parents {
		if val.Parents != nil {
			getAnc(val.Parents, ancestors)
		}
	}
}

func (f *familyTree) GetDescendents(id int) (map[int]*Node, error) {
	if f.nodes == nil {
		return nil, fmt.Errorf("no nodes exists in graph ")
	}

	_, exists := f.nodes[id]
	if !exists {
		return nil, fmt.Errorf("no such nodes exists")
	}

	descendents := make(map[int]*Node)
	fmt.Println(f.nodes[id].Children)
	getDes(f.nodes[id].Children, descendents)
	return descendents, nil
}

func getDes(children map[int]*Node, descendents map[int]*Node) {
	for _, val := range children {
		_, exists := descendents[val.Id]
		if !exists {
			descendents[val.Id] = val
			fmt.Println(val.Id)
		}
	}
	for _, val := range children {
		if val.Children != nil {
			getDes(val.Children, descendents)
		}
	}
}
