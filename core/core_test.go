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

	result := Any(numbers, isFive)

	assert.True(t, result)
}

func Test_AnyIntsReturnsFalse(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isNegativeOne := func(x int) bool {
		return x == -1
	}

	result := Any(numbers, isNegativeOne)

	assert.False(t, result)
}

// ALL
func Test_AllIntsReturnsTrue(t *testing.T) {
	numbers := []int{0, 2, 4, 6, 8}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := All(numbers, isEven)

	assert.True(t, result)
}

func Test_AllIntsReturnsFalse(t *testing.T) {
	numbers := []int{0, 2, 4, 6, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := All(numbers, isEven)

	assert.False(t, result)
}

// WHERE
func Test_WhereIntsReturnsEven(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := Where(numbers, isEven)

	assert.Equal(t, result, []int{0, 2, 4, 6, 8})
}

func Test_WhereIntsReturnsOdd(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := Where(numbers, isOdd)

	assert.Equal(t, result, []int{1, 3, 5, 7, 9})
}

// FIRST
func Test_FirstIntsReturnsEven(t *testing.T) {
	numbers := []int{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := First(numbers, isEven, -1)

	assert.Nil(t, err)
	assert.Equal(t, result, 0)
}

func Test_FirstIntsReturnsNotFound(t *testing.T) {
	numbers := []int{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := First(numbers, isOdd, -1)

	assert.NotNil(t, err)
	assert.Equal(t, err, errors.NotFoundError{})
	assert.Equal(t, result, -1)
}
