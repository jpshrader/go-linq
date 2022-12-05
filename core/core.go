package core

type WhereFilter[T any] func(x T) bool

func Where[T any](filter WhereFilter[T], src []T) (ret []T) {
	for _, item := range src {
		if filter(item) {
			ret = append(ret, item)
		}
	}
	return
}