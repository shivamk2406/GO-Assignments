package graph

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
	return &familyTree{nodes: make(map[int]*Node)}
}

type Node struct {
	Id       int
	Name     string
	Parents  map[int]*Node
	Children map[int]*Node
}

func (f *familyTree) AddNode(id int, name string) error {
	if _, exists := f.nodes[id]; exists {
		return fmt.Errorf(SingleNodeString, id, NodeAlreadyExistsErr)
	}

	f.nodes[id] = &Node{Id: id, Name: name}
	return nil
}

func (f *familyTree) AddEdge(id1 int, id2 int) error {
	if len(f.nodes) == 0 {
		return GraphEmptyErr
	}

	if _, exists := f.nodes[id1]; !exists {
		return fmt.Errorf(SingleNodeString, id1, NodeDNEErr)
	}

	if _, exists := f.nodes[id2]; !exists {
		return fmt.Errorf(SingleNodeString, id1, NodeDNEErr)
	}
	if _, exists := f.nodes[id1].Children[id2]; exists {
		return fmt.Errorf(DoubleNodeString, id1, id2, DependencyAlreadyExists)
	}

	id1Ancestors, _ := f.GetAncestors(id1)

	if _, exists := id1Ancestors[id2]; exists {
		return fmt.Errorf(DoubleNodeString, id1, id2, CyclicDependencyErr)
	}

	id2Descendants, _ := f.GetDescendents(id2)

	if _, exists := id2Descendants[id2]; exists {
		return fmt.Errorf(DoubleNodeString, id1, id2, CyclicDependencyErr)
	}

	if f.nodes[id1].Children == nil {
		f.nodes[id1].Children = make(map[int]*Node)
	}
	f.nodes[id1].Children[id2] = f.nodes[id2]

	if f.nodes[id2].Parents == nil {
		f.nodes[id2].Parents = make(map[int]*Node)
	}
	f.nodes[id2].Parents[id1] = f.nodes[id1]

	return nil
}

func (f *familyTree) DeleteNode(id int) error {
	if len(f.nodes) == 0 {
		return GraphEmptyErr
	}

	if _, exists := f.nodes[id]; !exists {
		return fmt.Errorf(SingleNodeString, id, NodeDNEErr)
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
	if len(f.nodes) == 0 {
		return GraphEmptyErr
	}

	if _, exists := f.nodes[id1]; !exists {
		return fmt.Errorf(SingleNodeString, id1, NodeDNEErr)
	}

	if _, exists := f.nodes[id2]; !exists {
		return fmt.Errorf(SingleNodeString, id1, NodeDNEErr)
	}

	firstNodeChildren := f.nodes[id1].Children
	if firstNodeChildren == nil {
		return fmt.Errorf(SingleNodeString, id1, NoChildrenExistErr)
	}

	if _, exists := firstNodeChildren[id2]; !exists {
		return fmt.Errorf(DoubleNodeString, id1, id2, NoSuchDependencyErr)
	}

	secondNodeParents := f.nodes[id2].Parents
	if secondNodeParents == nil {
		return fmt.Errorf(SingleNodeString, id2, NoParentsExistErr)
	}

	if _, exists := secondNodeParents[id1]; !exists {
		return fmt.Errorf(DoubleNodeString, id1, id2, NoSuchDependencyErr)
	}

	delete(secondNodeParents, id1)
	delete(firstNodeChildren, id2)

	return nil
}

func (f *familyTree) GetChildren(id int) (map[int]*Node, error) {
	if len(f.nodes) == 0 {
		return nil, GraphEmptyErr
	}

	if _, exists := f.nodes[id]; !exists {
		return nil, fmt.Errorf(SingleNodeString, id, NodeDNEErr)
	}

	if f.nodes[id].Children == nil {
		return nil, fmt.Errorf(SingleNodeString, id, NoChildrenExistErr)
	}

	return f.nodes[id].Children, nil
}

func (f *familyTree) GetParents(id int) (map[int]*Node, error) {
	if len(f.nodes) == 0 {
		return nil, GraphEmptyErr
	}

	if _, exists := f.nodes[id]; !exists {
		return nil, fmt.Errorf(SingleNodeString, id, NodeDNEErr)
	}

	if f.nodes[id].Parents == nil {
		return nil, fmt.Errorf(SingleNodeString, id, NoParentsExistErr)
	}

	return f.nodes[id].Parents, nil
}

func (f *familyTree) GetAncestors(id int) (map[int]*Node, error) {
	if len(f.nodes) == 0 {
		return nil, GraphEmptyErr
	}

	if _, exists := f.nodes[id]; !exists {
		return nil, fmt.Errorf(SingleNodeString, id, NodeDNEErr)
	}

	ancestors := make(map[int]*Node)
	getAnc(f.nodes[id], ancestors)

	if len(ancestors) == 0 {
		return nil, fmt.Errorf(SingleNodeString, id, NoAncestorsExistErr)
	}
	return ancestors, nil
}

func getAnc(node *Node, ancestors map[int]*Node) {
	for _, val := range node.Parents {
		if _, exists := ancestors[val.Id]; !exists {
			ancestors[val.Id] = val
		}

	}

	for _, val := range node.Parents {
		if val != nil {
			getAnc(val, ancestors)
		}
	}
}

func (f *familyTree) GetDescendents(id int) (map[int]*Node, error) {
	if len(f.nodes) == 0 {
		return nil, GraphEmptyErr
	}

	if _, exists := f.nodes[id]; !exists {
		return nil, fmt.Errorf(SingleNodeString, id, NodeDNEErr)
	}

	descendents := make(map[int]*Node)
	getDes(f.nodes[id], descendents)

	if len(descendents) == 0 {
		return nil, fmt.Errorf(SingleNodeString, id, NoDescendantsExistErr)
	}

	return descendents, nil
}

func getDes(node *Node, descendents map[int]*Node) {
	for _, val := range node.Children {
		if _, exists := descendents[val.Id]; !exists {
			descendents[val.Id] = val
		}
	}

	for _, val := range node.Children {
		if val != nil {
			getDes(val, descendents)
		}
	}
}
