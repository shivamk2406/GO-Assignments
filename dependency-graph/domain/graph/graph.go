package graph

import (
	"github.com/pkg/errors"
)

type FamilyTree interface {
	AddNode(int, string) error
	AddEdge(int, int) error
	DeleteNode(int) error
	DeleteEdge(int, int) error
	GetParents(int) (map[int]*Node, error)
	GetChildren(int) (map[int]*Node, error)
	GetAncestors(int, map[int]*Node) error
	GetDescendents(int, map[int]*Node) error
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
		return errors.Wrapf(NodeAlreadyExistsErr, "id : %d", id)
	}

	f.nodes[id] = &Node{Id: id, Name: name}
	return nil
}

func (f *familyTree) AddEdge(id1 int, id2 int) error {
	if len(f.nodes) == 0 {
		return GraphEmptyErr
	}

	if _, exists := f.nodes[id1]; !exists {
		return errors.Wrapf(NodeDNEErr, "id: %d", id1)
	}

	if _, exists := f.nodes[id2]; !exists {
		return errors.Wrapf(NodeDNEErr, "id: %d", id2)
	}
	if _, exists := f.nodes[id1].Children[id2]; exists {
		return errors.Wrapf(DependencyAlreadyExists, "id: %d and id : %d", id1, id2)
	}
	id1Ancestors := make(map[int]*Node)
	err := f.GetAncestors(id1, id1Ancestors)
	if err != nil {
		return err
	}

	if _, exists := id1Ancestors[id2]; exists {
		return errors.Wrapf(CyclicDependencyErr, "id: %d and id : %d", id1, id2)
	}

	id2Descendants := make(map[int]*Node)
	err = f.GetDescendents(id2, id2Descendants)
	if err != nil {
		if err != nil {
			return err
		}
	}

	if _, exists := id2Descendants[id2]; exists {
		return errors.Wrapf(CyclicDependencyErr, "id: %d and id : %d", id1, id2)
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
		return errors.Wrapf(NodeDNEErr, "id : %d", id)
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
		return errors.Wrapf(NodeDNEErr, "id: %d", id1)
	}

	if _, exists := f.nodes[id2]; !exists {
		return errors.Wrapf(NodeDNEErr, "id: %d", id2)
	}

	firstNodeChildren := f.nodes[id1].Children
	if firstNodeChildren == nil {
		return errors.Wrapf(NoChildrenExistErr, "id: %d", id1)
	}

	if _, exists := firstNodeChildren[id2]; !exists {
		return errors.Wrapf(NoSuchDependencyErr, "id: %d and id : %d", id1, id2)
	}

	secondNodeParents := f.nodes[id2].Parents
	if secondNodeParents == nil {
		return errors.Wrapf(NoParentsExistErr, "id: %d", id2)
	}

	if _, exists := secondNodeParents[id1]; !exists {
		return errors.Wrapf(NoSuchDependencyErr, "id: %d and id : %d", id1, id2)
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
		return nil, errors.Wrapf(NodeDNEErr, "id : %d", id)
	}

	if f.nodes[id].Children == nil {
		return nil, errors.Wrapf(NoChildrenExistErr, "id: %d", id)
	}

	return f.nodes[id].Children, nil
}

func (f *familyTree) GetParents(id int) (map[int]*Node, error) {
	if len(f.nodes) == 0 {
		return nil, GraphEmptyErr
	}

	if _, exists := f.nodes[id]; !exists {
		return nil, errors.Wrapf(NodeDNEErr, "id : %d", id)
	}

	if f.nodes[id].Parents == nil {
		return nil, errors.Wrapf(NoParentsExistErr, "id: %d", id)
	}

	return f.nodes[id].Parents, nil
}

func (f *familyTree) GetAncestors(id int, ancestors map[int]*Node) error {
	if _, exists := f.nodes[id]; !exists {
		return errors.Wrapf(NodeDNEErr, "id : %d", id)
	}

	for _, parent := range f.nodes[id].Parents {
		if _, exists := ancestors[parent.Id]; !exists {
			ancestors[parent.Id] = parent
			err := f.GetAncestors(parent.Id, ancestors)
			if err != nil {
				return errors.Wrapf(NoAncestorsExistErr, "id : %d", id)
			}
		}
	}
	return nil
}

func (f *familyTree) GetDescendents(id int, descendants map[int]*Node) error {
	if _, exists := f.nodes[id]; !exists {
		return errors.Wrapf(NodeDNEErr, "id : %d", id)
	}

	for _, child := range f.nodes[id].Children {
		if _, exists := descendants[child.Id]; !exists {
			descendants[child.Id] = child
			err := f.GetDescendents(child.Id, descendants)
			if err != nil {
				return errors.Wrapf(NoDescendantsExistErr, "id: %d", id)
			}
		}
	}
	return nil
}
