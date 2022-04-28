package graph

import "github.com/pkg/errors"

const (
	SingleNodeString = "id %d: %w"
	DoubleNodeString = "id %d & id %d : %w "
)

var (
	GraphEmptyErr           = errors.Errorf("Graph Empty")
	NodeDNEErr              = errors.Errorf("Node Not Found")
	NodeAlreadyExistsErr    = errors.Errorf("Node Already Exists")
	DependencyAlreadyExists = errors.Errorf("Dependency Already Exists")
	NoSuchDependencyErr     = errors.Errorf("No such Dependency Exists")
	CyclicDependencyErr     = errors.Errorf("Cyclic Dependency Exits")
	NoParentsExistErr       = errors.Errorf("No parents exists for the node")
	NoChildrenExistErr      = errors.Errorf("No Children Exists for the node")
	NoAncestorsExistErr     = errors.Errorf("No ancestors exists for the node ")
	NoDescendantsExistErr   = errors.Errorf("No Descendants exists for the node")
)
