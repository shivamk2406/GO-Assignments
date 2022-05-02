package view

type Choice int

const (
	ImmediateParentsChoice Choice = iota + 1
	ImmediateChildrenChoice
	AncestorsChoice
	DescendentsChoice
	DeleteDependencyChoice
	DeleteNodeChoice
	AddDependencyChoice
	AddNodeChoice
	ExitChoice
)
