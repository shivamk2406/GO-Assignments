package node

type FamilyTree interface {
	AddChildren()
	AddParents()
	GetParents() map[int]Node
	GetChildren() map[int]Node
}

type Node struct {
	Id       int
	Name     string
	Parents  map[int]Node
	Children map[int]Node
}

func (parent *Node) AddChildren(child Node) {
	parent.Children[child.Id] = child
}

func (child *Node) AddParents(parent Node) {
	child.Parents[parent.Id] = parent
}
