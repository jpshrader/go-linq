package receiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slice(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := AsSlice(numbers).
		Skip(5).
		Take(5).
		Where(isEven)

	assert.Equal(t, Slice[int]{6, 8}, result)
}
