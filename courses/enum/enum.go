package enum

type Course int

//go:generate  -type=Course

const (
	a Course = iota
	b
	c
	d
	e
	f
)
