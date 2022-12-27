package golinq

import (
	"github.com/jpshrader/go-linq/core"
)

type Slice[T any] []T

func AsSlice[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}

func (slice Slice[T]) Any(filter core.Matcher[T]) bool {
	return core.Any(slice, filter)
}

func (slice Slice[T]) All(filter core.Matcher[T]) bool {
	return core.All(slice, filter)
}

func (slice Slice[T]) Where(filter core.Matcher[T]) Slice[T] {
	return core.Where(slice, filter)
}

func (slice Slice[T]) First(filter core.Matcher[T], defaultVal T) (T, error) {
	return core.First(slice, filter, defaultVal)
}

type Map[K comparable, V any] map[K]V

func AsMap[K comparable, V any](m map[K]V) Map[K, V] {
	return Map[K, V](m)
}
