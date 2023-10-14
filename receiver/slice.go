package receiver

import (
	golinq "github.com/jpshrader/go-linq"
)

type Slice[T any] []T

func AsSlice[T any](s []T) Slice[T] {
	return Slice[T](s)
}

func (s Slice[T]) Append(item T) Slice[T] {
	return golinq.Append(s, item)
}

func (s Slice[T]) Concat(items []T) Slice[T] {
	return golinq.Concat(s, items)
}

func (s Slice[T]) Any(filter func(x T) bool) bool {
	return golinq.Any(s, filter)
}

func (s Slice[T]) All(filter func(x T) bool) bool {
	return golinq.All(s, filter)
}

func (s Slice[T]) IsEmpty() bool {
	return golinq.IsEmpty(s)
}

func (s Slice[T]) Count() int {
	return golinq.Count(s)
}

func (s Slice[T]) Contains(isEqual func(x, y T) bool, item T) bool {
	return golinq.Contains(s, isEqual, item)
}

func (s Slice[T]) Where(filter func(x T) bool) Slice[T] {
	return golinq.Where(s, filter)
}

func (s Slice[T]) Except(filter func(x T) bool) Slice[T] {
	return golinq.Except(s, filter)
}

func (s Slice[T]) First(filter func(x T) bool) (T, error) {
	return golinq.First(s, filter)
}

func (s Slice[T]) Last(filter func(x T) bool) (T, error) {
	return golinq.Last(s, filter)
}

func (s Slice[T]) Take(count int) Slice[T] {
	return golinq.Take(s, count)
}

func (s Slice[T]) Skip(count int) Slice[T] {
	return golinq.Skip(s, count)
}

func (s Slice[T]) OrderBy(comparer func(x, y int) bool) Slice[T] {
	return golinq.OrderBy(s, comparer)
}
