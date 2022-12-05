package golinq

import (
	"github.com/jpshrader/go-linq/core"
)

type Slice[T any] []T

func AsSlice[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}

func (slice Slice[T]) Any(filter core.Matcher[T]) bool {
	return core.Any(filter, slice)
}

func (slice Slice[T]) All(filter core.Matcher[T]) bool {
	return core.All(filter, slice)
}

func (slice Slice[T]) Where(filter core.Matcher[T]) Slice[T] {
	return core.Where(filter, slice)
}

func (slice Slice[T]) First(filter core.Matcher[T], defaultVal T) (T, error) {
	return core.First(filter, defaultVal, slice)
}
