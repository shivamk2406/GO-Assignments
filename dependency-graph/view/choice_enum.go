package view

type Choice int

const (
	GetImmediateParents Choice = iota + 1
	GetImmediateChildren
	GetAncestors
	GetDescendents
	DeleteDependency
	DeleteNode
	AddDependency
	AddNode
	Exit
)
