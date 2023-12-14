package golinq

import (
    "sort"

    "golang.org/x/exp/constraints"
)

func Append[T any](items []T, item T) []T {
    return append(items, item)
}

func Any[T any](items []T, isMatch func(x T) bool) bool {
    for _, item := range items {
        if isMatch(item) {
            return true
        }
    }
    return false
}

func All[T any](items []T, isMatch func(x T) bool) bool {
    for _, item := range items {
        if !isMatch(item) {
            return false
        }
    }
    return true
}

func Concat[T any](items []T, newItems []T) []T {
    return append(items, newItems...)
}

func IsEmpty[T any](items []T) bool {
    return len(items) == 0
}

func Count[T any](items []T) int {
    return len(items)
}

func Contains[T any](items []T, isEqual func(x, y T) bool, item T) bool {
    for _, element := range items {
        if isEqual(item, element) {
            return true
        }
    }
    return false
}

func Where[T any](items []T, isMatch func(x T) bool) []T {
    out := make([]T, 0, len(items))
    for _, item := range items {
        if isMatch(item) {
            out = append(out, item)
        }
    }
    return out
}

func Except[T any](items []T, isMatch func(x T) bool) []T {
    out := make([]T, 0, len(items))
    for _, item := range items {
        if !isMatch(item) {
            out = append(out, item)
        }
    }
    return out
}

func First[T any](items []T, isMatch func(x T) bool) (T, error) {
    if isMatch == nil {
        isMatch = NilPredicate[T]
    }
    for _, item := range items {
        if isMatch(item) {
            return item, nil
        }
    }
    return *new(T), NotFoundError{}
}

func Last[T any](items []T, isMatch func(x T) bool) (T, error) {
    var last T
    found := false
    for _, item := range items {
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

func Take[T any](items []T, count int) []T {
    out := make([]T, 0, count)
    for i, item := range items {
        if i >= count {
            return out
        }
        out = append(out, item)
    }
    return out
}

func Skip[T any](items []T, count int) []T {
    if count >= len(items) {
        return []T{}
    }

    out := make([]T, 0, len(items)-count)
    for i := count; i < len(items); i++ {
        out = Append(out, items[i])
    }
    return out
}

func Select[T, R any](items []T, mapper func(x T) R) []R {
    out := make([]R, 0, len(items))
    for _, item := range items {
        out = append(out, mapper(item))
    }
    return out
}

func Distinct[T comparable](items []T) []T {
    out := make([]T, 0, len(items))
    existingItems := make(map[T]bool, len(items))
    for _, item := range items {
        if _, found := existingItems[item]; !found {
            out = append(out, item)
            existingItems[item] = true
        }
    }
    return out
}

func DistinctC[T any, R comparable](items []T, mapper func(x T) R) []T {
    out := make([]T, 0, len(items))
    existingItems := make(map[R]bool, len(items))
    for _, item := range items {
        itemId := mapper(item)
        if !existingItems[itemId] {
            out = append(out, item)
            existingItems[itemId] = true
        }
    }
    return out
}

func OrderBy[T any](items []T, comparer func(x, y int) bool) []T {
    sort.Slice(items, comparer)
    return items
}

func OrderByAscending[T constraints.Ordered](items []T) []T {
    comparer := func(i, j int) bool {
        return items[i] < items[j]
    }
    sort.Slice(items, comparer)
    return items
}

func OrderByDescending[T constraints.Ordered](items []T) []T {
    comparer := func(i, j int) bool {
        return items[i] > items[j]
    }
    sort.Slice(items, comparer)
    return items
}

func Sum[T Number](items []T) T {
    out := T(0)
    for _, item := range items {
        out += item
    }
    return out
}

func SumC[T any, R Number](items []T, mapper func(x T) R) R {
    out := R(0)
    for _, item := range items {
        out += mapper(item)
    }
    return out
}

func Average[T Number](items []T) float64 {
    length := len(items)
    if length > 0 {
        return float64(Sum(items)) / float64(length)
    }
    return 0
}

func AverageC[T any, R Number](items []T, mapper func(x T) R) float64 {
    length := len(items)
    if length > 0 {
        return float64(SumC(items, mapper)) / float64(length)
    }
    return 0
}

func Max[T Number](items []T) T {
    out := T(0)
    for _, item := range items {
        if item > out {
            out = item
        }
    }
    return out
}

func MaxC[T any, R constraints.Ordered](items []T, mapper func(x T) R) T {
    out := *new(T)
    for _, item := range items {
        if mapper(item) > mapper(out) {
            out = item
        }
    }
    return out
}

func Min[T Number](items []T) T {
    out := *new(T)
    for _, item := range items {
        if item < out {
            out = item
        }
    }
    return out
}

func MinC[T any, R constraints.Ordered](items []T, mapper func(x T) R) T {
    out := *new(T)
    for _, item := range items {
        if mapper(item) < mapper(out) {
            out = item
        }
    }
    return out
}

func Reverse[T any](items []T) []T {
    out := make([]T, 0, len(items))
    for i := len(items) - 1; i >= 0; i-- {
        out = append(out, items[i])
    }
    return out
}

func GroupBy[T any, R comparable](items []T, mapper func(x T) R) map[R][]T {
    out := make(map[R][]T, len(items))
    for _, item := range items {
        key := mapper(item)
        out[key] = append(out[key], item)
    }
    return out
}

func Aggregate[T any](items []T, seed T, accumulator func(T, T) T) T {
    for _, item := range items {
        seed = accumulator(seed, item)
    }
    return seed
}
