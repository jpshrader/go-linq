package receiver

import (
	golinq "github.com/jpshrader/go-linq"
)

type NumericSlice[T golinq.Number] []T

func AsNumericSlice[T golinq.Number](s []T) NumericSlice[T] {
	return NumericSlice[T](s)
}

func (s NumericSlice[T]) Append(item T) NumericSlice[T] {
	return golinq.Append(s, item)
}

func (s NumericSlice[T]) Concat(items NumericSlice[T]) NumericSlice[T] {
	return golinq.Concat(s, items)
}

func (s NumericSlice[T]) Any(filter func(x T) bool) bool {
	return golinq.Any(s, filter)
}

func (s NumericSlice[T]) All(filter func(x T) bool) bool {
	return golinq.All(s, filter)
}

func (s NumericSlice[T]) IsEmpty() bool {
	return golinq.IsEmpty(s)
}

func (s NumericSlice[T]) Count() int {
	return golinq.Count(s)
}

func (s NumericSlice[T]) Contains(item T) bool {
	isEqual := func(x, y T) bool {
		return x == y
	}

	return golinq.Contains(s, isEqual, item)
}

func (s NumericSlice[T]) Where(filter func(x T) bool) NumericSlice[T] {
	return golinq.Where(s, filter)
}

func (s NumericSlice[T]) Except(filter func(x T) bool) NumericSlice[T] {
	return golinq.Except(s, filter)
}

func (s NumericSlice[T]) First(filter func(x T) bool) (T, error) {
	return golinq.First(s, filter)
}

func (s NumericSlice[T]) Last(filter func(x T) bool) (T, error) {
	return golinq.Last(s, filter)
}

func (s NumericSlice[T]) Take(count int) NumericSlice[T] {
	return golinq.Take(s, count)
}

func (s NumericSlice[T]) Skip(count int) NumericSlice[T] {
	return golinq.Skip(s, count)
}

func (s NumericSlice[T]) OrderBy(comparer func(x, y int) bool) NumericSlice[T] {
	return golinq.OrderBy(s, comparer)
}

func (s NumericSlice[T]) OrderByAscending() NumericSlice[T] {
	return golinq.OrderByAscending(s)
}

func (s NumericSlice[T]) OrderByDescending() NumericSlice[T] {
	return golinq.OrderByDescending(s)
}

func (s NumericSlice[T]) Sum() T {
	return golinq.Sum(s)
}

func (s NumericSlice[T]) Average() float64 {
	return golinq.Average(s)
}

func (s NumericSlice[T]) Max() T {
	return golinq.Max(s)
}

func (s NumericSlice[T]) Min() T {
	return golinq.Min(s)
}

func (s NumericSlice[T]) Reverse() NumericSlice[T] {
	return golinq.Reverse(s)
}

func (s NumericSlice[T]) Distinct() NumericSlice[T] {
	return golinq.Distinct(s)
}
