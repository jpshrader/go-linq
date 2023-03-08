package golinq

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Append[T any](src []T, item T) []T {
	return append(src, item)
}

func Any[T any](src []T, isMatch func(x T) bool) bool {
	for _, item := range src {
		if isMatch(item) {
			return true
		}
	}
	return false
}

func All[T any](src []T, isMatch func(x T) bool) bool {
	for _, item := range src {
		if !isMatch(item) {
			return false
		}
	}
	return true
}

func Concat[T any](src []T, items []T) []T {
	return append(src, items...)
}

func IsEmpty[T any](src []T) bool {
	return len(src) == 0
}

func Count[T any](src []T) int {
	return len(src)
}

func Contains[T any](src []T, isEqual func(x, y T) bool, item T) bool {
	for _, element := range src {
		if isEqual(item, element) {
			return true
		}
	}
	return false
}

func Where[T any](src []T, isMatch func(x T) bool) (ret []T) {
	for _, item := range src {
		if isMatch(item) {
			ret = append(ret, item)
		}
	}
	return
}

func Except[T any](src []T, isMatch func(x T) bool) (ret []T) {
	for _, item := range src {
		if !isMatch(item) {
			ret = append(ret, item)
		}
	}
	return
}

func First[T any](src []T, isMatch func(x T) bool) (T, error) {
	if isMatch == nil {
		isMatch = NilPredicate[T]
	}
	for _, item := range src {
		if isMatch(item) {
			return item, nil
		}
	}
	return *new(T), NotFoundError{}
}

func Last[T any](src []T, isMatch func(x T) bool) (T, error) {
	var last T
	found := false
	for _, item := range src {
		if isMatch(item) {
			last = item
			found = true
		}
	}

	if found {
		return last, nil
	}
	return *new(T), NotFoundError{}
}

func Take[T any](src []T, count int) (ret []T) {
	for i, item := range src {
		if i < count {
			ret = append(ret, item)
		} else {
			return
		}
	}
	return
}

func Skip[T any](src []T, count int) (ret []T) {
	if count > len(src) {
		return
	}

	for i := count; i < len(src); i++ {
		ret = append(ret, src[i])
	}
	return
}

func Select[T, R any](src []T, mapper func(x T) R) (ret []R) {
	for _, item := range src {
		ret = append(ret, mapper(item))
	}
	return
}

func Distinct[T comparable](src []T) (ret []T) {
	existingItems := make(map[T]bool, 0)
	for _, item := range src {
		if !existingItems[item] {
			ret = append(ret, item)
			existingItems[item] = true
		}
	}
	return
}

func DistinctC[T any, R comparable](src []T, mapper func(x T) R) (ret []T) {
	existingItems := make(map[R]bool)
	for _, item := range src {
		itemId := mapper(item)
		if !existingItems[itemId] {
			ret = append(ret, item)
			existingItems[itemId] = true
		}
	}
	return
}

func OrderBy[T any](src []T, comparer func(x, y int) bool) []T {
	sort.Slice(src, comparer)
	return src
}

func OrderByAscending[T constraints.Ordered](src []T) []T {
	comparer := func(i, j int) bool {
		return src[i] < src[j]
	}
	sort.Slice(src, comparer)
	return src
}

func OrderByDescending[T constraints.Ordered](src []T) []T {
	comparer := func(i, j int) bool {
		return src[i] > src[j]
	}
	sort.Slice(src, comparer)
	return src
}

func Sum[T Number](src []T) (ret T) {
	for _, item := range src {
		ret += item
	}
	return
}

func SumC[T any, R Number](src []T, mapper func(x T) R) (ret R) {
	for _, item := range src {
		ret += mapper(item)
	}
	return
}

func Average[T Number](src []T) float64 {
	length := len(src)
	if length > 0 {
		return float64(Sum(src)) / float64(length)
	}
	return 0
}

func AverageC[T any, R Number](src []T, mapper func(x T) R) float64 {
	length := len(src)
	if length > 0 {
		return float64(SumC(src, mapper)) / float64(length)
	}
	return 0
}

func Max[T Number](src []T) (ret T) {
	for _, item := range src {
		if item > ret {
			ret = item
		}
	}
	return
}

func MaxC[T any, R constraints.Ordered](src []T, mapper func(x T) R) (ret T) {
	for _, item := range src {
		if mapper(item) > mapper(ret) {
			ret = item
		}
	}
	return
}

func Min[T Number](src []T) (ret T) {
	for _, item := range src {
		if item < ret {
			ret = item
		}
	}
	return
}

func MinC[T any, R constraints.Ordered](src []T, mapper func(x T) R) (ret T) {
	for _, item := range src {
		if mapper(item) < mapper(ret) {
			ret = item
		}
	}
	return
}

func Reverse[T any](src []T) (ret []T) {
	for i := len(src) - 1; i >= 0; i-- {
		ret = append(ret, src[i])
	}
	return
}

func GroupBy[T any, R comparable](src []T, mapper func(x T) R) (ret map[R][]T) {
	ret = make(map[R][]T)
	for _, item := range src {
		key := mapper(item)
		ret[key] = append(ret[key], item)
	}
	return
}

func Aggregate[T any](src []T, seed T, accumulator func(T, T) T) T {
	for _, item := range src {
		seed = accumulator(seed, item)
	}
	return seed
}
