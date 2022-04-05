package item

type ItemType int

const (
	raw ItemType = iota
	manufactured
	imported
)
