package golinq

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func NilPredicate[T any](x T) bool {
	return true
}