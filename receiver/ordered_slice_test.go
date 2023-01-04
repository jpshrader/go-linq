package receiver

import (
	"testing"

	golinq "github.com/jpshrader/go-linq"
	"github.com/stretchr/testify/assert"
)

// AS ORDERED SLICE
func Test_AsOrderedSlice(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4}

	result := AsOrderedSlice(numbers)

	assert.Equal(t, OrderedSlice[int]{0, 1, 2, 3, 4}, result)
}

// ANY
func Test_OrderedAnyIntsReturnsTrue(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isFive := func(x int) bool {
		return x == 5
	}

	result := numbers.Any(isFive)

	assert.True(t, result)
}

func Test_OrderedAnyIntsReturnsFalse(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isNegativeOne := func(x int) bool {
		return x == -1
	}

	result := numbers.Any(isNegativeOne)

	assert.False(t, result)
}

// ALL
func Test_OrderedAllIntsReturnsTrue(t *testing.T) {
	numbers := OrderedSlice[int]{0, 2, 4, 6, 8}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.All(isEven)

	assert.True(t, result)
}

func Test_OrderedAllIntsReturnsFalse(t *testing.T) {
	numbers := OrderedSlice[int]{0, 2, 4, 6, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.All(isEven)

	assert.False(t, result)
}

// WHERE
func Test_OrderedWhereIntsReturnsEven(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.Where(isEven)

	assert.Equal(t, result, OrderedSlice[int]{0, 2, 4, 6, 8})
}

func Test_OrderedWhereIntsReturnsOdd(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := numbers.Where(isOdd)

	assert.Equal(t, result, OrderedSlice[int]{1, 3, 5, 7, 9})
}

// EXCEPT
func Test_OrderedExceptIntsReturnsNotEven(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.Except(isEven)

	assert.Equal(t, OrderedSlice[int]{1, 3, 5, 7, 9}, result)
}

func Test_OrderedExceptIntsReturnsNotOdd(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := numbers.Except(isOdd)

	assert.Equal(t, OrderedSlice[int]{0, 2, 4, 6, 8}, result)
}

// FIRST
func Test_OrderedFirstIntsReturnsEven(t *testing.T) {
	numbers := OrderedSlice[int]{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := numbers.First(isEven)

	assert.Nil(t, err)
	assert.Equal(t, result, 0)
}

func Test_OrderedFirstIntsReturnsNotFound(t *testing.T) {
	numbers := OrderedSlice[int]{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := numbers.First(isOdd)

	assert.NotNil(t, err)
	assert.Equal(t, err, golinq.NotFoundError{})
	assert.Equal(t, result, 0)
}

// LAST
func Test_OrderedLastIntsReturnsEven(t *testing.T) {
	numbers := OrderedSlice[int]{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := numbers.Last(isEven)

	assert.Nil(t, err)
	assert.Equal(t, result, 2)
}

func Test_OrderedLastIntsReturnsNotFound(t *testing.T) {
	numbers := OrderedSlice[int]{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := numbers.Last(isOdd)

	assert.NotNil(t, err)
	assert.Equal(t, err, golinq.NotFoundError{})
	assert.Equal(t, result, 0)
}

// TAKE
func Test_OrderedTakeIntsReturnsFirst5(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Take(5)

	assert.Equal(t, OrderedSlice[int]{0, 1, 2, 3, 4}, result)
}

func Test_OrderedTakeIntsReturnsAllAvailable(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1}

	result := numbers.Take(5)

	assert.Equal(t, OrderedSlice[int]{0, 1}, result)
}

func Test_OrderedTakeIntsReturnsAllAvailableNil(t *testing.T) {
	numbers := OrderedSlice[int]{}

	result := numbers.Take(5)

	var expected OrderedSlice[int]
	assert.Equal(t, expected, result)
}

// SKIP
func Test_OrderedSkipIntsReturnsExceptFirst5(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Skip(5)

	assert.Equal(t, OrderedSlice[int]{5, 6, 7, 8, 9}, result)
}

func Test_OrderedSkipIntsReturnsAllAvailableExceptFirst5(t *testing.T) {
	numbers := OrderedSlice[int]{0, 1}

	result := numbers.Skip(5)

	var expected OrderedSlice[int]
	assert.Equal(t, expected, result)
}

// ORDER BY
func Test_OrderedOrderByIntsReturnsAscending(t *testing.T) {
	numbers := OrderedSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	lessThan := func(i, j int) bool {
		return numbers[i] < numbers[j]
	}

	result := numbers.OrderBy(lessThan)

	assert.Equal(t, OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}

// ORDER BY ASCENDING
func Test_OrderedOrderByAscendingIntsReturns(t *testing.T) {
	numbers := OrderedSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := numbers.OrderByAscending()

	assert.Equal(t, OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}

// ORDER BY DESCENDING
func Test_OrderedOrderByDescendingIntsReturns(t *testing.T) {
	numbers := OrderedSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := numbers.OrderByDescending()

	assert.Equal(t, OrderedSlice[int]{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, result)
}

func Test_OrderedOrderByIntsReturnsDescending(t *testing.T) {
	numbers := Slice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	greaterThan := func(i, j int) bool {
		return numbers[i] > numbers[j]
	}

	result := numbers.OrderBy(greaterThan)

	assert.Equal(t, Slice[int]{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, result)
}