package receiver

import (
	golinq "github.com/jpshrader/go-linq"
	"golang.org/x/exp/constraints"
)

type OrderedSlice[T constraints.Ordered] []T

func AsOrderedSlice[T constraints.Ordered](slice []T) OrderedSlice[T] {
	return OrderedSlice[T](slice)
}

func (slice OrderedSlice[T]) Any(filter golinq.Predicate[T]) bool {
	return golinq.Any(slice, filter)
}

func (slice OrderedSlice[T]) All(filter golinq.Predicate[T]) bool {
	return golinq.All(slice, filter)
}

func (slice OrderedSlice[T]) IsEmpty() bool {
	return golinq.IsEmpty(slice)
}

func (slice OrderedSlice[T]) Count() int {
	return golinq.Count(slice)
}

func (slice OrderedSlice[T]) Contains(item T) bool {
	isEqual := func (x, y T) bool {
		return x == y
	}

	return golinq.Contains(slice, isEqual, item)
}

func (slice OrderedSlice[T]) Where(filter golinq.Predicate[T]) OrderedSlice[T] {
	return golinq.Where(slice, filter)
}

func (slice OrderedSlice[T]) Except(filter golinq.Predicate[T]) OrderedSlice[T] {
	return golinq.Except(slice, filter)
}

func (slice OrderedSlice[T]) First(filter golinq.Predicate[T]) (T, error) {
	return golinq.First(slice, filter)
}

func (slice OrderedSlice[T]) Last(filter golinq.Predicate[T]) (T, error) {
	return golinq.Last(slice, filter)
}

func (slice OrderedSlice[T]) Take(count int) OrderedSlice[T] {
	return golinq.Take(slice, count)
}

func (slice OrderedSlice[T]) Skip(count int) OrderedSlice[T] {
	return golinq.Skip(slice, count)
}

func (slice OrderedSlice[T]) OrderBy(comparer golinq.IndexComparer[T]) OrderedSlice[T] {
	return golinq.OrderBy(slice, comparer)
}

func (slice OrderedSlice[T]) OrderByAscending() OrderedSlice[T] {
	return golinq.OrderByAscending(slice)
}

func (slice OrderedSlice[T]) OrderByDescending() OrderedSlice[T] {
	return golinq.OrderByDescending(slice)
}
