package core

import (
	"sort"

	"github.com/jpshrader/go-linq/errors"
)

type Matcher[T any] func(x T) bool

type Transform[T, R any] func(x T) R

type Compare[T any] func(x, y int) bool

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Any[T any](src []T, isMatch Matcher[T]) bool {
	for _, item := range src {
		if isMatch(item) {
			return true
		}
	}
	return false
}

func All[T any](src []T, isMatch Matcher[T]) bool {
	for _, item := range src {
		if !isMatch(item) {
			return false
		}
	}
	return true
}

func Where[T any](src []T, isMatch Matcher[T]) (ret []T) {
	for _, item := range src {
		if isMatch(item) {
			ret = append(ret, item)
		}
	}
	return
}

func Except[T any](src []T, isMatch Matcher[T]) (ret []T) {
	for _, item := range src {
		if !isMatch(item) {
			ret = append(ret, item)
		}
	}
	return
}

func First[T any](src []T, isMatch Matcher[T], defaultVal T) (T, error) {
	for _, item := range src {
		if isMatch(item) {
			return item, nil
		}
	}
	return defaultVal, errors.NotFoundError{}
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

func Select[T, R any](src []T, transformer Transform[T, R]) (ret []R) {
	for _, item := range src {
		ret = append(ret, transformer(item))
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

func DistinctStruct[T any, R comparable](src []T, transformer Transform[T, R]) (ret []T) {
	existingItems := make(map[R]bool)
	for _, item := range src {
		itemId := transformer(item)
		if !existingItems[itemId] {
			ret = append(ret, item)
			existingItems[itemId] = true
		}
	}
	return
}

func OrderBy[T any](src []T, comparer Compare[T]) []T {
	sort.Slice(src, comparer)
	return src
}

func Sum[T Number](src []T) (ret T) {
	for _, item := range src {
		ret += item
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

func Max[T Number](src []T) (ret T) {
	for _, item := range src {
		if item > ret {
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
