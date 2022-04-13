package courses

type CourseType int

//go:generate  -type=CourseType

const (
	a CourseType = iota
	b
	c
	d
	e
	f
)
