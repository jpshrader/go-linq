package receiver

import golinq "github.com/jpshrader/go-linq"

type NumericSlice[T golinq.Number] []T

func AsNumericSlice[T golinq.Number](slice []T) NumericSlice[T] {
	return NumericSlice[T](slice)
}

func (slice NumericSlice[T]) Append(item T) NumericSlice[T] {
	return golinq.Append(slice, item)
}

func (slice NumericSlice[T]) Concat(items NumericSlice[T]) NumericSlice[T] {
	return golinq.Concat(slice, items)
}

func (slice NumericSlice[T]) Any(filter func(x T) bool) bool {
	return golinq.Any(slice, filter)
}

func (slice NumericSlice[T]) All(filter func(x T) bool) bool {
	return golinq.All(slice, filter)
}

func (slice NumericSlice[T]) IsEmpty() bool {
	return golinq.IsEmpty(slice)
}

func (slice NumericSlice[T]) Count() int {
	return golinq.Count(slice)
}

func (slice NumericSlice[T]) Contains(item T) bool {
	isEqual := func(x, y T) bool {
		return x == y
	}

	return golinq.Contains(slice, isEqual, item)
}

func (slice NumericSlice[T]) Where(filter func(x T) bool) NumericSlice[T] {
	return golinq.Where(slice, filter)
}

func (slice NumericSlice[T]) Except(filter func(x T) bool) NumericSlice[T] {
	return golinq.Except(slice, filter)
}

func (slice NumericSlice[T]) First(filter func(x T) bool) (T, error) {
	return golinq.First(slice, filter)
}

func (slice NumericSlice[T]) Last(filter func(x T) bool) (T, error) {
	return golinq.Last(slice, filter)
}

func (slice NumericSlice[T]) Take(count int) NumericSlice[T] {
	return golinq.Take(slice, count)
}

func (slice NumericSlice[T]) Skip(count int) NumericSlice[T] {
	return golinq.Skip(slice, count)
}

func (slice NumericSlice[T]) OrderBy(comparer func(x, y int) bool) NumericSlice[T] {
	return golinq.OrderBy(slice, comparer)
}

func (slice NumericSlice[T]) OrderByAscending() NumericSlice[T] {
	return golinq.OrderByAscending(slice)
}

func (slice NumericSlice[T]) OrderByDescending() NumericSlice[T] {
	return golinq.OrderByDescending(slice)
}

func (slice NumericSlice[T]) Sum() T {
	return golinq.Sum(slice)
}

func (slice NumericSlice[T]) Average() float64 {
	return golinq.Average(slice)
}

func (slice NumericSlice[T]) Max() T {
	return golinq.Max(slice)
}

func (slice NumericSlice[T]) Min() T {
	return golinq.Min(slice)
}

func (slice NumericSlice[T]) Reverse() NumericSlice[T] {
	return golinq.Reverse(slice)
}

func (slice NumericSlice[T]) Distinct() NumericSlice[T] {
	return golinq.Distinct(slice)
}
