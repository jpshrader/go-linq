package receiver

import "github.com/jpshrader/go-linq/linq"

type NumericSlice[T linq.Number] []T

func AsNumericSlice[T linq.Number](slice []T) NumericSlice[T] {
	return NumericSlice[T](slice)
}

func (slice NumericSlice[T]) Append(item T) NumericSlice[T] {
	return linq.Append(slice, item)
}

func (slice NumericSlice[T]) Concat(items []T) NumericSlice[T] {
	return linq.Concat(slice, items)
}

func (slice NumericSlice[T]) Any(filter func(x T) bool) bool {
	return linq.Any(slice, filter)
}

func (slice NumericSlice[T]) All(filter func(x T) bool) bool {
	return linq.All(slice, filter)
}

func (slice NumericSlice[T]) IsEmpty() bool {
	return linq.IsEmpty(slice)
}

func (slice NumericSlice[T]) Count() int {
	return linq.Count(slice)
}

func (slice NumericSlice[T]) Contains(item T) bool {
	isEqual := func(x, y T) bool {
		return x == y
	}

	return linq.Contains(slice, isEqual, item)
}

func (slice NumericSlice[T]) Where(filter func(x T) bool) NumericSlice[T] {
	return linq.Where(slice, filter)
}

func (slice NumericSlice[T]) Except(filter func(x T) bool) NumericSlice[T] {
	return linq.Except(slice, filter)
}

func (slice NumericSlice[T]) First(filter func(x T) bool) (T, error) {
	return linq.First(slice, filter)
}

func (slice NumericSlice[T]) Last(filter func(x T) bool) (T, error) {
	return linq.Last(slice, filter)
}

func (slice NumericSlice[T]) Take(count int) NumericSlice[T] {
	return linq.Take(slice, count)
}

func (slice NumericSlice[T]) Skip(count int) NumericSlice[T] {
	return linq.Skip(slice, count)
}

func (slice NumericSlice[T]) OrderBy(comparer func(x, y int) bool) NumericSlice[T] {
	return linq.OrderBy(slice, comparer)
}

func (slice NumericSlice[T]) Sum() T {
	return linq.Sum(slice)
}

func (slice NumericSlice[T]) Average() float64 {
	return linq.Average(slice)
}

func (slice NumericSlice[T]) Max() T {
	return linq.Max(slice)
}

func (slice NumericSlice[T]) Min() T {
	return linq.Min(slice)
}
