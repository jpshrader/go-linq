package receiver

import (
	golinq "github.com/jpshrader/go-linq"
)

type Slice[T any] []T

func AsSlice[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}

func (slice Slice[T]) Append(item T) Slice[T] {
	return golinq.Append(slice, item)
}

func (slice Slice[T]) Concat(items []T) Slice[T] {
	return golinq.Concat(slice, items)
}

func (slice Slice[T]) Any(filter func(x T) bool) bool {
	return golinq.Any(slice, filter)
}

func (slice Slice[T]) All(filter func(x T) bool) bool {
	return golinq.All(slice, filter)
}

func (slice Slice[T]) IsEmpty() bool {
	return golinq.IsEmpty(slice)
}

func (slice Slice[T]) Count() int {
	return golinq.Count(slice)
}

func (slice Slice[T]) Contains(isEqual func(x, y T) bool, item T) bool {
	return golinq.Contains(slice, isEqual, item)
}

func (slice Slice[T]) Where(filter func(x T) bool) Slice[T] {
	return golinq.Where(slice, filter)
}

func (slice Slice[T]) Except(filter func(x T) bool) Slice[T] {
	return golinq.Except(slice, filter)
}

func (slice Slice[T]) First(filter func(x T) bool) (T, error) {
	return golinq.First(slice, filter)
}

func (slice Slice[T]) Last(filter func(x T) bool) (T, error) {
	return golinq.Last(slice, filter)
}

func (slice Slice[T]) Take(count int) Slice[T] {
	return golinq.Take(slice, count)
}

func (slice Slice[T]) Skip(count int) Slice[T] {
	return golinq.Skip(slice, count)
}

func (slice Slice[T]) OrderBy(comparer func(x, y int) bool) Slice[T] {
	return golinq.OrderBy(slice, comparer)
}
