package graph

import "github.com/pkg/errors"

var (
	GraphEmptyErr           = errors.New("Graph Empty")
	NodeDNEErr              = errors.New("Node Not Found")
	NodeAlreadyExistsErr    = errors.New("Node Already Exists")
	DependencyAlreadyExists = errors.New("Dependency Already Exists")
	NoSuchDependencyErr     = errors.New("No such Dependency Exists")
	CyclicDependencyErr     = errors.New("Cyclic Dependency Exits")
	NoParentsExistErr       = errors.New("No parents exists for the node")
	NoChildrenExistErr      = errors.New("No Children Exists for the node")
	NoAncestorsExistErr     = errors.New("No ancestors exists for the node ")
	NoDescendantsExistErr   = errors.New("No Descendants exists for the node")
)
