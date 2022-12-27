package core

import (
	"sort"

	"github.com/jpshrader/go-linq/errors"

	"golang.org/x/exp/constraints"
)

type Matcher[T any] func(x T) bool

type Transformer[T, R any] func(x T) R

type Comparer[T any] func(x, y int) bool

type Number interface {
	constraints.Integer | constraints.Float
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

func Last[T any](src []T, isMatch Matcher[T], defaultVal T) (T, error) {
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

func Select[T, R any](src []T, transformer Transformer[T, R]) (ret []R) {
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

func DistinctC[T any, R comparable](src []T, transformer Transformer[T, R]) (ret []T) {
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

func OrderBy[T any](src []T, comparer Comparer[T]) []T {
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

func SumC[T any, R Number](src []T, transformer Transformer[T, R]) (ret R) {
	for _, item := range src {
		ret += transformer(item)
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

func AverageC[T any, R Number](src []T, transformer Transformer[T, R]) float64 {
	length := len(src)
	if length > 0 {
		return float64(SumC(src, transformer)) / float64(length)
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

func MaxC[T any, R constraints.Ordered](src []T, transformer Transformer[T, R]) (ret T) {
	for _, item := range src {
		if transformer(item) > transformer(ret) {
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

func MinC[T any, R constraints.Ordered](src []T, transformer Transformer[T, R]) (ret T) {
	for _, item := range src {
		if transformer(item) < transformer(ret) {
			ret = item
		}
	}
	return
}
