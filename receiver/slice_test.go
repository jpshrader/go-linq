package receiver

import (
	"testing"

	golinq "github.com/jpshrader/go-linq"
	"github.com/stretchr/testify/assert"
)

// AS X SLICE
func Test_AsSlice(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4}

	result := AsSlice(numbers)

	assert.Equal(t, Slice[int]{0, 1, 2, 3, 4}, result)
}

func Test_AsNumericSlice(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4}

	result := AsNumericSlice(numbers)

	assert.Equal(t, NumericSlice[int]{0, 1, 2, 3, 4}, result)
}

func Test_AsOrderedSlice(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4}

	result := AsOrderedSlice(numbers)

	assert.Equal(t, OrderedSlice[int]{0, 1, 2, 3, 4}, result)
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

// ORDER BY ASCENDING
func Test_OrderByAscendingIntsReturns(t *testing.T) {
	numbers := OrderedSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := numbers.OrderByAscending()

	assert.Equal(t, OrderedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}

// ORDER BY DESCENDING
func Test_OrderByDescendingIntsReturns(t *testing.T) {
	numbers := OrderedSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := numbers.OrderByDescending()

	assert.Equal(t, OrderedSlice[int]{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, result)
}

func Test_OrderByIntsReturnsDescending(t *testing.T) {
	numbers := Slice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	greaterThan := func(i, j int) bool {
		return numbers[i] > numbers[j]
	}

	result := numbers.OrderBy(greaterThan)

	assert.Equal(t, Slice[int]{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, result)
}

// MAX
func Test_MaxIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := numbers.Max()

	assert.Equal(t, 9, result)
}

func Test_MaxIntsReturnsNil(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Max()

	assert.Equal(t, 0, result)
}

// MIN
func Test_MinIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := numbers.Min()

	assert.Equal(t, 0, result)
}

func Test_MinIntsReturnsLast(t *testing.T) {
	numbers := NumericSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5, -1}

	result := numbers.Min()

	assert.Equal(t, -1, result)
}

func Test_MinIntsReturnsNil(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Min()

	assert.Equal(t, 0, result)
}

// SUM
func Test_SumIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Sum()

	assert.Equal(t, 45, result)
}

func Test_SumIntsReturnsNil(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Sum()

	assert.Equal(t, 0, result)
}

// AVERAGE
func Test_AverageIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Average()

	assert.Equal(t, 4.5, result)
}

func Test_AverageIntsReturnsNil(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Average()

	assert.Equal(t, float64(0), result)
}
