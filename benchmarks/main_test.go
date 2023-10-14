package main

import (
	"testing"

	golinq "github.com/jpshrader/go-linq"
	"github.com/jpshrader/go-linq/receiver"
)

func isEven(x int) bool {
	return x%2 == 0
}

var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var numerics = receiver.AsNumericSlice(numbers)

func Benchmark_Where_ForLoop(b *testing.B) {
	var evenNums []int
	for i := 0; i < b.N; i++ {
		evenNums = []int{}
		for _, n := range numbers {
			if isEven(n) {
				evenNums = append(evenNums, n)
			}
		}
	}
}

func Benchmark_Where_Vanilla(b *testing.B) {
	for i := 0; i < b.N; i++ {
		golinq.Where(numbers, isEven)
	}
}

func Benchmark_Where_Numeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numerics.Where(isEven)
	}
}

func Benchmark_Append(b *testing.B) {
	out := make([]int, 0, len(numbers))
	for _, num := range numbers {
		out = append(out, num)
	}
}

func Benchmark_Append_Resize(b *testing.B) {
	out := []int{}
	for _, num := range numbers {
		out = append(out, num)
	}
}

func Benchmark_Index(b *testing.B) {
	out := make([]int, len(numbers))
	for i, num := range numbers {
		out[i] = num
	}
}
