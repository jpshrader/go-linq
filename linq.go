package golinq

import (
	"github.com/jpshrader/go-linq/core"
	"golang.org/x/exp/constraints"
)

type Slice[T any] []T

type OrderedSlice[T constraints.Ordered] []T

type NumericSlice[T core.Number] []T

type ComparableSlice[T comparable] []T

type TransformableSlice[T, R any] []T

type ComparableTransformableSlice[T any, R comparable] []T

func AsSlice[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}

func AsNumericSlice[T core.Number](slice []T) NumericSlice[T] {
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

func (slice Slice[T]) Any(filter core.Matcher[T]) bool {
	return core.Any(slice, filter)
}

func (slice Slice[T]) All(filter core.Matcher[T]) bool {
	return core.All(slice, filter)
}

func (slice Slice[T]) Where(filter core.Matcher[T]) Slice[T] {
	return core.Where(slice, filter)
}

func (slice Slice[T]) Except(filter core.Matcher[T]) Slice[T] {
	return core.Except(slice, filter)
}

func (slice Slice[T]) First(filter core.Matcher[T], defaultVal T) (T, error) {
	return core.First(slice, filter, defaultVal)
}

func (slice Slice[T]) Take(count int) Slice[T] {
	return core.Take(slice, count)
}

func (slice Slice[T]) Skip(count int) Slice[T] {
	return core.Skip(slice, count)
}

func (slice TransformableSlice[T, R]) Select(transformer core.Transformer[T, R]) Slice[R] {
	return core.Select(slice, transformer)
}

func (slice ComparableSlice[T]) Distinct() Slice[T] {
	return core.Distinct(slice)
}

func (slice ComparableTransformableSlice[T, R]) DistinctStruct(transformer core.Transformer[T, R]) Slice[T] {
	return core.DistinctStruct(slice, transformer)
}

func (slice Slice[T]) OrderBy(comparer core.Comparer[T]) Slice[T] {
	return core.OrderBy(slice, comparer)
}

func (slice OrderedSlice[T]) OrderByAscending() OrderedSlice[T] {
	return core.OrderByAscending(slice)
}

func (slice OrderedSlice[T]) OrderByDescending() OrderedSlice[T] {
	return core.OrderByDescending(slice)
}

func (slice NumericSlice[T]) Sum() T {
	return core.Sum(slice)
}

func (slice NumericSlice[T]) Average() float64 {
	return core.Average(slice)
}

func (slice NumericSlice[T]) Max() T {
	return core.Max(slice)
}

func (slice NumericSlice[T]) Min() T {
	return core.Min(slice)
}
