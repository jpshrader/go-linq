package slices

import (
	golinq "github.com/jpshrader/go-linq"
	"golang.org/x/exp/constraints"
)

type OrderedSlice[T constraints.Ordered] []T

func AsOrderedSlice[T constraints.Ordered](s []T) OrderedSlice[T] {
	return OrderedSlice[T](s)
}

func (s OrderedSlice[T]) Append(item T) OrderedSlice[T] {
	return golinq.Append(s, item)
}

func (s OrderedSlice[T]) Concat(items []T) OrderedSlice[T] {
	return golinq.Concat(s, items)
}

func (s OrderedSlice[T]) Any(filter func(x T) bool) bool {
	return golinq.Any(s, filter)
}

func (s OrderedSlice[T]) All(filter func(x T) bool) bool {
	return golinq.All(s, filter)
}

func (s OrderedSlice[T]) IsEmpty() bool {
	return golinq.IsEmpty(s)
}

func (s OrderedSlice[T]) Count() int {
	return golinq.Count(s)
}

func (s OrderedSlice[T]) Contains(item T) bool {
	isEqual := func(x, y T) bool {
		return x == y
	}

	return golinq.Contains(s, isEqual, item)
}

func (s OrderedSlice[T]) Where(filter func(x T) bool) OrderedSlice[T] {
	return golinq.Where(s, filter)
}

func (s OrderedSlice[T]) Except(filter func(x T) bool) OrderedSlice[T] {
	return golinq.Except(s, filter)
}

func (s OrderedSlice[T]) First(filter func(x T) bool) (T, error) {
	return golinq.First(s, filter)
}

func (s OrderedSlice[T]) Last(filter func(x T) bool) (T, error) {
	return golinq.Last(s, filter)
}

func (s OrderedSlice[T]) Take(count int) OrderedSlice[T] {
	return golinq.Take(s, count)
}

func (s OrderedSlice[T]) Skip(count int) OrderedSlice[T] {
	return golinq.Skip(s, count)
}

func (s OrderedSlice[T]) OrderBy(comparer func(x, y int) bool) OrderedSlice[T] {
	return golinq.OrderBy(s, comparer)
}

func (s OrderedSlice[T]) OrderByAscending() OrderedSlice[T] {
	return golinq.OrderByAscending(s)
}

func (s OrderedSlice[T]) OrderByDescending() OrderedSlice[T] {
	return golinq.OrderByDescending(s)
}
