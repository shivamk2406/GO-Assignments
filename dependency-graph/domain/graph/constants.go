package graph

const (
	GraphEmptyErr        = "No nodes exists in the graph!!! Graph empty"
	NodeDoNotExistErr    = "No such node with id: %d exists in the graph"
	NodeAlreadyExistsErr = "Node with the id: %d alredy exists"
	NoSuchDependencyErr  = "no relationship found between the nodes %d %d"
	CyclicDependencyErr  = "A relationship already exists between the nodes %d and %d"
)
