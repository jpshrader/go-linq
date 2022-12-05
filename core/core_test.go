package core

import (
	"testing"

	"github.com/jpshrader/go-linq/errors"
	"github.com/stretchr/testify/assert"
)

// ANY
func Test_AnyIntsReturnsTrue(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isFive := func(x int) bool {
		return x == 5
	}

	result := Any(isFive, numbers)

	assert.True(t, result)
}

func Test_AnyIntsReturnsFalse(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isNegativeOne := func(x int) bool {
		return x == -1
	}

	result := Any(isNegativeOne, numbers)

	assert.False(t, result)
}

// ALL
func Test_AllIntsReturnsTrue(t *testing.T) {
	numbers := []int{0, 2, 4, 6, 8}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := All(isEven, numbers)

	assert.True(t, result)
}

func Test_AllIntsReturnsFalse(t *testing.T) {
	numbers := []int{0, 2, 4, 6, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := All(isEven, numbers)

	assert.False(t, result)
}

// WHERE
func Test_WhereIntsReturnsEven(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := Where(isEven, numbers)

	assert.Equal(t, result, []int{0, 2, 4, 6, 8})
}

func Test_WhereIntsReturnsOdd(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := Where(isOdd, numbers)

	assert.Equal(t, result, []int{1, 3, 5, 7, 9})
}

// FIRST
func Test_FirstIntsReturnsEven(t *testing.T) {
	numbers := []int{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := First(isEven, -1, numbers)

	assert.Nil(t, err)
	assert.Equal(t, result, 0)
}

func Test_FirstIntsReturnsNotFound(t *testing.T) {
	numbers := []int{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := First(isOdd, -1, numbers)

	assert.NotNil(t, err)
	assert.Equal(t, err, errors.NotFoundError{})
	assert.Equal(t, result, -1)
}
