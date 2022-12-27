package receiver

import (
	"github.com/jpshrader/go-linq"
	"golang.org/x/exp/constraints"
)

type Slice[T any] []T

type OrderedSlice[T constraints.Ordered] []T

type NumericSlice[T golinq.Number] []T

type ComparableSlice[T comparable] []T

type TransformableSlice[T, R any] []T

type ComparableTransformableSlice[T any, R comparable] []T

func AsSlice[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}

func AsNumericSlice[T golinq.Number](slice []T) NumericSlice[T] {
	return NumericSlice[T](slice)
}

func AsComparableSlice[T comparable](slice []T) ComparableSlice[T] {
	return ComparableSlice[T](slice)
}

func AsTransformableSlice[T, R any](slice []T) TransformableSlice[T, R] {
	return TransformableSlice[T, R](slice)
}

func AsComparableTransformableSlice[T any, R comparable](slice []T) ComparableTransformableSlice[T, R] {
	return ComparableTransformableSlice[T, R](slice)
}

func (slice Slice[T]) Any(filter golinq.Matcher[T]) bool {
	return golinq.Any(slice, filter)
}

func (slice Slice[T]) All(filter golinq.Matcher[T]) bool {
	return golinq.All(slice, filter)
}

func (slice Slice[T]) Where(filter golinq.Matcher[T]) Slice[T] {
	return golinq.Where(slice, filter)
}

func (slice Slice[T]) Except(filter golinq.Matcher[T]) Slice[T] {
	return golinq.Except(slice, filter)
}

func (slice Slice[T]) First(filter golinq.Matcher[T]) (T, error) {
	return golinq.First(slice, filter)
}

func (slice Slice[T]) Last(filter golinq.Matcher[T]) (T, error) {
	return golinq.Last(slice, filter)
}

func (slice Slice[T]) Take(count int) Slice[T] {
	return golinq.Take(slice, count)
}

func (slice Slice[T]) Skip(count int) Slice[T] {
	return golinq.Skip(slice, count)
}

func (slice TransformableSlice[T, R]) Select(transformer golinq.Transformer[T, R]) Slice[R] {
	return golinq.Select(slice, transformer)
}

func (slice ComparableSlice[T]) Distinct() Slice[T] {
	return golinq.Distinct(slice)
}

func (slice ComparableTransformableSlice[T, R]) DistinctStruct(transformer golinq.Transformer[T, R]) Slice[T] {
	return golinq.DistinctC(slice, transformer)
}

func (slice Slice[T]) OrderBy(comparer golinq.Comparer[T]) Slice[T] {
	return golinq.OrderBy(slice, comparer)
}

func (slice OrderedSlice[T]) OrderByAscending() OrderedSlice[T] {
	return golinq.OrderByAscending(slice)
}

func (slice OrderedSlice[T]) OrderByDescending() OrderedSlice[T] {
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