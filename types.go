package golinq

import "golang.org/x/exp/constraints"

// INTERFACES
type Number interface {
	constraints.Integer | constraints.Float
}

// FUNCTIONS
type Predicate[T any] func(x T) bool

func EmptyPredicate[T any](x T) bool {
	return true
}

type Mapper[T, R any] func(x T) R

type Comparer[T any] func(x, y int) bool

// ERRORS
type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "no matching element(s) found"
}
