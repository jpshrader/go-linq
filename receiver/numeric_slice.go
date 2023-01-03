package receiver

import golinq "github.com/jpshrader/go-linq"

type NumericSlice[T golinq.Number] []T

func AsNumericSlice[T golinq.Number](slice []T) NumericSlice[T] {
	return NumericSlice[T](slice)
}

func (slice NumericSlice[T]) Any(filter golinq.Predicate[T]) bool {
	return golinq.Any(slice, filter)
}

func (slice NumericSlice[T]) All(filter golinq.Predicate[T]) bool {
	return golinq.All(slice, filter)
}

func (slice NumericSlice[T]) Where(filter golinq.Predicate[T]) NumericSlice[T] {
	return golinq.Where(slice, filter)
}

func (slice NumericSlice[T]) Except(filter golinq.Predicate[T]) NumericSlice[T] {
	return golinq.Except(slice, filter)
}

func (slice NumericSlice[T]) First(filter golinq.Predicate[T]) (T, error) {
	return golinq.First(slice, filter)
}

func (slice NumericSlice[T]) Last(filter golinq.Predicate[T]) (T, error) {
	return golinq.Last(slice, filter)
}

func (slice NumericSlice[T]) Take(count int) NumericSlice[T] {
	return golinq.Take(slice, count)
}

func (slice NumericSlice[T]) Skip(count int) NumericSlice[T] {
	return golinq.Skip(slice, count)
}

func (slice NumericSlice[T]) OrderBy(comparer golinq.Comparer[T]) NumericSlice[T] {
	return golinq.OrderBy(slice, comparer)
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