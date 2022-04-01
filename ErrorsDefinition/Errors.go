package errorsdefinition

type error interface {
	Error() string
}
type MyError struct{}
