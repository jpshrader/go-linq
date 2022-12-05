package core

import (
	"github.com/jpshrader/go-linq/errors"
)

type Matcher[T any] func(x T) bool

func Any[T any](isMatch Matcher[T], src []T) bool {
	for _, item := range src {
		if isMatch(item) {
			return true
		}
	}
	return false
}

func All[T any](isMatch Matcher[T], src []T) bool {
	for _, item := range src {
		if !isMatch(item) {
			return false
		}
	}
	return true
}

func Where[T any](isMatch Matcher[T], src []T) (ret []T) {
	for _, item := range src {
		if isMatch(item) {
			ret = append(ret, item)
		}
	}
	return
}

func First[T any](isMatch Matcher[T], defaultVal T, src []T) (T, error) {
	for _, item := range src {
		if isMatch(item) {
			return item, nil
		}
	}
	return defaultVal, errors.NotFoundError{}
}
