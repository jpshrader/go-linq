package receiver

import (
	"github.com/jpshrader/go-linq/linq"
	"golang.org/x/exp/constraints"
)

type OrderedSlice[T constraints.Ordered] []T

func AsOrderedSlice[T constraints.Ordered](slice []T) OrderedSlice[T] {
	return OrderedSlice[T](slice)
}

func (slice OrderedSlice[T]) Append(item T) OrderedSlice[T] {
	return linq.Append(slice, item)
}

func (slice OrderedSlice[T]) Concat(items []T) OrderedSlice[T] {
	return linq.Concat(slice, items)
}

func (slice OrderedSlice[T]) Any(filter func(x T) bool) bool {
	return linq.Any(slice, filter)
}

func (slice OrderedSlice[T]) All(filter func(x T) bool) bool {
	return linq.All(slice, filter)
}

func (slice OrderedSlice[T]) IsEmpty() bool {
	return linq.IsEmpty(slice)
}

func (slice OrderedSlice[T]) Count() int {
	return linq.Count(slice)
}

func (slice OrderedSlice[T]) Contains(item T) bool {
	isEqual := func(x, y T) bool {
		return x == y
	}

	return linq.Contains(slice, isEqual, item)
}

func (slice OrderedSlice[T]) Where(filter func(x T) bool) OrderedSlice[T] {
	return linq.Where(slice, filter)
}

func (slice OrderedSlice[T]) Except(filter func(x T) bool) OrderedSlice[T] {
	return linq.Except(slice, filter)
}

func (slice OrderedSlice[T]) First(filter func(x T) bool) (T, error) {
	return linq.First(slice, filter)
}

func (slice OrderedSlice[T]) Last(filter func(x T) bool) (T, error) {
	return linq.Last(slice, filter)
}

func (slice OrderedSlice[T]) Take(count int) OrderedSlice[T] {
	return linq.Take(slice, count)
}

func (slice OrderedSlice[T]) Skip(count int) OrderedSlice[T] {
	return linq.Skip(slice, count)
}

func (slice OrderedSlice[T]) OrderBy(comparer func(x, y int) bool) OrderedSlice[T] {
	return linq.OrderBy(slice, comparer)
}

func (slice OrderedSlice[T]) OrderByAscending() OrderedSlice[T] {
	return linq.OrderByAscending(slice)
}

func (slice OrderedSlice[T]) OrderByDescending() OrderedSlice[T] {
	return linq.OrderByDescending(slice)
}
