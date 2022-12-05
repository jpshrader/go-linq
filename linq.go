package golinq

type Slice[T any] []T

func (slice Slice[T]) Where(filter WhereFilter[T]) (ret Slice[T]) {
	return Where(filter, slice)
}