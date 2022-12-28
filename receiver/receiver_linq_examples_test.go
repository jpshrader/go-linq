package receiver

import (
	"testing"

	golinq "github.com/jpshrader/go-linq"
	"github.com/stretchr/testify/assert"
)

func Test_Slice(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := AsSlice(numbers).
		Skip(5).
		Take(5).
		Where(isEven).
		First(golinq.EmptyPredicate[int])

	assert.Nil(t, err)
	assert.Equal(t, 6, result)
}
