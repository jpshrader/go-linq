package receiver

import (
	"testing"

	golinq "github.com/jpshrader/go-linq"
	"github.com/stretchr/testify/assert"
)

// AS SLICE
func Test_AsSlice(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4}

	result := AsSlice(numbers)

	assert.Equal(t, Slice[int]{0, 1, 2, 3, 4}, result)
}

// ANY
func Test_AnyIntsReturnsTrue(t *testing.T) {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isFive := func(x int) bool {
		return x == 5
	}

	result := numbers.Any(isFive)

	assert.True(t, result)
}

func Test_AnyIntsReturnsFalse(t *testing.T) {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isNegativeOne := func(x int) bool {
		return x == -1
	}

	result := numbers.Any(isNegativeOne)

	assert.False(t, result)
}

// ALL
func Test_AllIntsReturnsTrue(t *testing.T) {
	numbers := Slice[int]{0, 2, 4, 6, 8}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.All(isEven)

	assert.True(t, result)
}

func Test_AllIntsReturnsFalse(t *testing.T) {
	numbers := Slice[int]{0, 2, 4, 6, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.All(isEven)

	assert.False(t, result)
}

// WHERE
func Test_WhereIntsReturnsEven(t *testing.T) {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.Where(isEven)

	assert.Equal(t, result, Slice[int]{0, 2, 4, 6, 8})
}

func Test_WhereIntsReturnsOdd(t *testing.T) {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := numbers.Where(isOdd)

	assert.Equal(t, result, Slice[int]{1, 3, 5, 7, 9})
}

// EXCEPT
func Test_ExceptIntsReturnsNotEven(t *testing.T) {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.Except(isEven)

	assert.Equal(t, Slice[int]{1, 3, 5, 7, 9}, result)
}

func Test_ExceptIntsReturnsNotOdd(t *testing.T) {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := numbers.Except(isOdd)

	assert.Equal(t, Slice[int]{0, 2, 4, 6, 8}, result)
}

// FIRST
func Test_FirstIntsReturnsEven(t *testing.T) {
	numbers := Slice[int]{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := numbers.First(isEven)

	assert.Nil(t, err)
	assert.Equal(t, result, 0)
}

func Test_FirstIntsReturnsNotFound(t *testing.T) {
	numbers := Slice[int]{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := numbers.First(isOdd)

	assert.NotNil(t, err)
	assert.Equal(t, err, golinq.NotFoundError{})
	assert.Equal(t, result, 0)
}

// LAST
func Test_LastIntsReturnsEven(t *testing.T) {
	numbers := Slice[int]{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := numbers.Last(isEven)

	assert.Nil(t, err)
	assert.Equal(t, result, 2)
}

func Test_LastIntsReturnsNotFound(t *testing.T) {
	numbers := Slice[int]{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := numbers.Last(isOdd)

	assert.NotNil(t, err)
	assert.Equal(t, err, golinq.NotFoundError{})
	assert.Equal(t, result, 0)
}

// TAKE
func Test_TakeIntsReturnsFirst5(t *testing.T) {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Take(5)

	assert.Equal(t, Slice[int]{0, 1, 2, 3, 4}, result)
}

func Test_TakeIntsReturnsAllAvailable(t *testing.T) {
	numbers := Slice[int]{0, 1}

	result := numbers.Take(5)

	assert.Equal(t, Slice[int]{0, 1}, result)
}

func Test_TakeIntsReturnsAllAvailableNil(t *testing.T) {
	numbers := Slice[int]{}

	result := numbers.Take(5)

	var expected Slice[int]
	assert.Equal(t, expected, result)
}

// SKIP
func Test_SkipIntsReturnsExceptFirst5(t *testing.T) {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Skip(5)

	assert.Equal(t, Slice[int]{5, 6, 7, 8, 9}, result)
}

func Test_SkipIntsReturnsAllAvailableExceptFirst5(t *testing.T) {
	numbers := Slice[int]{0, 1}

	result := numbers.Skip(5)

	var expected Slice[int]
	assert.Equal(t, expected, result)
}

// ORDER BY
func Test_OrderByIntsReturnsAscending(t *testing.T) {
	numbers := Slice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	lessThan := func(i, j int) bool {
		return numbers[i] < numbers[j]
	}

	result := numbers.OrderBy(lessThan)

	assert.Equal(t, Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}