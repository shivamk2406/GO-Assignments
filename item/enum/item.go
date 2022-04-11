package enum

//go:generate  -type=ItemType
type ItemType int

const (
	Raw ItemType = iota
	Manufactured
	Imported
)
