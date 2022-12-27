package core

import (
	"sort"

	"github.com/jpshrader/go-linq/errors"
)

type Matcher[T any] func(x T) bool

type Transform[T, R any] func(x T) R

type Compare[T any] func(x, y int) bool

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

func First[T any](src []T, isMatch Matcher[T], defaultVal T) (T, error) {
	for _, item := range src {
		if isMatch(item) {
			return item, nil
		}
	}
	return defaultVal, errors.NotFoundError{}
}

func Select[T, R any](src []T, transformer Transform[T, R]) (ret []R) {
	for _, item := range src {
		ret = append(ret, transformer(item))
	}
	return
} 

func Distinct[T any, R comparable](src []T, transformer Transform[T, R]) (ret []T) {
	existingItems := make(map[R]bool)
	for _, item := range src {
		itemId := transformer(item)
		if existingItems[itemId] {
			ret = append(ret, item)
			existingItems[itemId] = true
		}
	}
	return
}

func OrderBy[T any](src []T, comparer Compare[T]) ([]T) {
	sort.Slice(src, comparer)
	return src
}