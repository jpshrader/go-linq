# Linq with Go

An attempt to recreate the ease and usefulness of the [linq dotnet library](https://learn.microsoft.com/en-us/dotnet/api/system.linq.enumerable?view=net-7.0) in golang.

## Usage

Real-world examples can be found [here](cmd/main.go)

```
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
```

## Helpful Info

1. https://go.dev/blog/intro-generics
2. https://gosamples.dev/generics-cheatsheet/

## Setup

This project utilises Visual Studio Dev Containers to manage dependencies. For more information on VS Code Development Containers, see the following:

1. https://code.visualstudio.com/docs/remote/containers
2. https://code.visualstudio.com/learn/develop-cloud/containers
3. https://github.com/microsoft/vscode-remote-try-go
