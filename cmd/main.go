package main

import (
	"fmt"

	golinq "github.com/jpshrader/go-linq"
	"github.com/jpshrader/go-linq/receiver"
)

func isEven(x int) bool {
    return x%2 == 0
}

func isOdd(x int) bool {
    return x%2 == 1
}

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("numbers: %v\n", numbers)

	evenNums := golinq.Where(numbers, isEven)
	fmt.Printf("even numbers: %v\n", evenNums)

	nums := receiver.AsNumericSlice(numbers)
	oddNums := nums.Where(isOdd)
	fmt.Printf("odd numbers: %v\n", oddNums)
}