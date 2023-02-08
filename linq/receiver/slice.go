package receiver

import "github.com/jpshrader/go-linq/linq"

type Slice[T any] []T

func AsSlice[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}

func (slice Slice[T]) Append(item T) Slice[T] {
	return linq.Append(slice, item)
}

func (slice Slice[T]) Concat(items []T) Slice[T] {
	return linq.Concat(slice, items)
}

func (slice Slice[T]) Any(filter func(x T) bool) bool {
	return linq.Any(slice, filter)
}

func (slice Slice[T]) All(filter func(x T) bool) bool {
	return linq.All(slice, filter)
}

func (slice Slice[T]) IsEmpty() bool {
	return linq.IsEmpty(slice)
}

func (slice Slice[T]) Count() int {
	return linq.Count(slice)
}

func (slice Slice[T]) Contains(isEqual func(x, y T) bool, item T) bool {
	return linq.Contains(slice, isEqual, item)
}

func (slice Slice[T]) Where(filter func(x T) bool) Slice[T] {
	return linq.Where(slice, filter)
}

func (slice Slice[T]) Except(filter func(x T) bool) Slice[T] {
	return linq.Except(slice, filter)
}

func (slice Slice[T]) First(filter func(x T) bool) (T, error) {
	return linq.First(slice, filter)
}

func (slice Slice[T]) Last(filter func(x T) bool) (T, error) {
	return linq.Last(slice, filter)
}

func (slice Slice[T]) Take(count int) Slice[T] {
	return linq.Take(slice, count)
}

func (slice Slice[T]) Skip(count int) Slice[T] {
	return linq.Skip(slice, count)
}

func (slice Slice[T]) OrderBy(comparer func(x, y int) bool) Slice[T] {
	return linq.OrderBy(slice, comparer)
}
