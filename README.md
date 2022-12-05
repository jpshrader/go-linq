# Linq with Go

## Usage

```
import (
    "fmt"

    "github.com/jpshrader/go-linq"
)

func isEven(x int) bool {
    return x % 2 == 0
}

func isOdd(x int) bool {
    return !isEven(x)
}

func main() {
    numbers := golinq.AsSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

    evenNumbers := numbers.Where(isEven)
    oddNumbers := numbers.Where(isOdd)

    fmt.Printf("even numbers: %v", evenNumbers)
    fmt.Printf("odd numbers:  %v", oddNumbers)
}
```

## Setup

This project utilises Visual Studio Dev Containers to manage dependencies. For more information on VS Code Development Containers, see the following:

1. https://code.visualstudio.com/docs/remote/containers
2. https://code.visualstudio.com/learn/develop-cloud/containers
3. https://github.com/microsoft/vscode-remote-try-go
