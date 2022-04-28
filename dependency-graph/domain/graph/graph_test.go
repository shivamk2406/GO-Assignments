package graph

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type TestScenario struct {
	description string
	nodeerr     error
	nodeID      int
}

func populateGraph(familyTree FamilyTree) {
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

func TestGetParents(t *testing.T) {
	graph := NewFamilyTree()
	populateGraph(graph)
	scenarios := []TestScenario{
		{
			description: "all nodes are fine",
			nodeerr:     nil,
			nodeID:      3,
		},
		{
			description: "parents for node do not exists",
			nodeerr:     NoParentsExistErr,
			nodeID:      1,
		},
	}

	for _, scenario := range scenarios {
		_, err := graph.GetParents(scenario.nodeID)
		require.Equal(t, scenario.nodeerr, errors.Unwrap(err))
	}

}

func TestGetChildren(t *testing.T) {
	graph := NewFamilyTree()
	populateGraph(graph)
	scenarios := []TestScenario{
		{
			description: "children exists for the node",
			nodeID:      1,
			nodeerr:     nil,
		},
		{
			description: "children do not exists for the node",
			nodeID:      6,
			nodeerr:     NoChildrenExistErr,
		},
	}

	for _, scenario := range scenarios {
		_, err := graph.GetChildren(scenario.nodeID)
		require.Equal(t, scenario.nodeerr, errors.Unwrap(err))
	}

}

func TestGetAncestors(t *testing.T) {
	graph := NewFamilyTree()
	populateGraph(graph)
	scenarios := []TestScenario{
		{
			description: "ancestors exists for the node",
			nodeerr:     nil,
			nodeID:      4,
		},
		{
			description: "no ancestor exists for the node",
			nodeerr:     NoAncestorsExistErr,
			nodeID:      1,
		},
	}
	for _, scenario := range scenarios {
		_, err := graph.GetAncestors(scenario.nodeID)
		require.Equal(t, scenario.nodeerr, errors.Unwrap(err))
	}

}

func TestGetDescendants(t *testing.T) {
	graph := NewFamilyTree()
	populateGraph(graph)
	scenarios := []TestScenario{
		{
			description: "descendants exists for the node",
			nodeerr:     nil,
			nodeID:      1,
		},
		{
			description: "no descendanst exists for the node",
			nodeerr:     NoDescendantsExistErr,
			nodeID:      6,
		},
	}
	for _, scenario := range scenarios {
		_, err := graph.GetDescendents(scenario.nodeID)
		require.Equal(t, scenario.nodeerr, errors.Unwrap(err))
	}

}

type TestCyclicScenario struct {
	description string
	id1         int
	id2         int
	cyclicError error
}

func TestCyclicDependency(t *testing.T) {
	graph := NewFamilyTree()
	populateGraph(graph)
	scenarios := []TestCyclicScenario{
		{
			description: "all nodes addition are fine",
			id1:         6,
			id2:         7,
			cyclicError: nil,
		},
		{
			description: "dependency already exits between the two nodes",
			id1:         3,
			id2:         1,
			cyclicError: CyclicDependencyErr,
		},
	}
	for _, scenario := range scenarios {
		err := graph.AddEdge(scenario.id1, scenario.id2)
		require.Equal(t, scenario.cyclicError, errors.Unwrap(err))
	}

}
