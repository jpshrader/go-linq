package linq

import "golang.org/x/exp/constraints"

// INTERFACES
type Number interface {
	constraints.Integer | constraints.Float
}

func NilPredicate[T any](x T) bool {
	return true
}

// ERRORS
type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "no matching element(s) found"
}
