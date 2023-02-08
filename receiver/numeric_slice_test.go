package receiver

import (
	"testing"

	golinq "github.com/jpshrader/go-linq"
	"github.com/stretchr/testify/assert"
)

// AS NUMERIC SLICE
func Test_AsNumericSlice(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4}

	result := AsNumericSlice(numbers)

	assert.Equal(t, NumericSlice[int]{0, 1, 2, 3, 4}, result)
}

// APPEND
func Test_NumericSliceAppendIntsReturns(t *testing.T) {
	src := NumericSlice[int]{1, 3, 5, 7, 9}

	result := src.Append(0)

	assert.Equal(t, NumericSlice[int]{1, 3, 5, 7, 9, 0}, result)
}

func Test_NumericSliceAppendAdditionalEmptyIntsReturns(t *testing.T) {
	var src NumericSlice[int]

	result := src.Append(0)

	assert.Equal(t, NumericSlice[int]{0}, result)
}

// CONCAT
func Test_NumericSliceConcatIntsReturns(t *testing.T) {
	src := NumericSlice[int]{1, 3, 5, 7, 9}
	additional := NumericSlice[int]{0, 2, 4, 6, 8}

	result := src.Concat(additional)

	assert.Equal(t, NumericSlice[int]{1, 3, 5, 7, 9, 0, 2, 4, 6, 8}, result)
}

func Test_NumericSliceConcatSrcEmptyIntsReturns(t *testing.T) {
	src := NumericSlice[int]{1, 3, 5, 7, 9}
	var additional NumericSlice[int]

	result := src.Concat(additional)

	assert.Equal(t, NumericSlice[int]{1, 3, 5, 7, 9}, result)
}

func Test_NumericSliceConcatAdditionalEmptyIntsReturns(t *testing.T) {
	var src NumericSlice[int]
	additional := NumericSlice[int]{0, 2, 4, 6, 8}

	result := src.Concat(additional)

	assert.Equal(t, NumericSlice[int]{0, 2, 4, 6, 8}, result)
}

// ANY
func Test_NumericAnyIntsReturnsTrue(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isFive := func(x int) bool {
		return x == 5
	}

	result := numbers.Any(isFive)

	assert.True(t, result)
}

func Test_NumericAnyIntsReturnsFalse(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isNegativeOne := func(x int) bool {
		return x == -1
	}

	result := numbers.Any(isNegativeOne)

	assert.False(t, result)
}

// ALL
func Test_NumericAllIntsReturnsTrue(t *testing.T) {
	numbers := NumericSlice[int]{0, 2, 4, 6, 8}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.All(isEven)

	assert.True(t, result)
}

func Test_NumericAllIntsReturnsFalse(t *testing.T) {
	numbers := NumericSlice[int]{0, 2, 4, 6, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.All(isEven)

	assert.False(t, result)
}

// ISEMPTY
func Test_NumericSliceIsEmptyIntsReturn(t *testing.T) {
	numbers := NumericSlice[int]{0, 2}

	result := numbers.IsEmpty()

	assert.False(t, result)
}

func Test_NumericSliceIsEmptyEmptyIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.IsEmpty()

	assert.True(t, result)
}

func Test_NumericSliceIsEmptyNilIntsReturns(t *testing.T) {
	var numbers NumericSlice[int]

	result := numbers.IsEmpty()

	assert.True(t, result)
}

// COUNT
func Test_NumericSliceCountIntsReturn(t *testing.T) {
	numbers := NumericSlice[int]{0, 2}

	result := numbers.Count()

	assert.Equal(t, 2, result)
}

func Test_NumericSliceCountEmptyIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Count()

	assert.Equal(t, 0, result)
}

func Test_NumericSliceCountNilIntsReturns(t *testing.T) {
	var numbers NumericSlice[int]

	result := numbers.Count()

	assert.Equal(t, 0, result)
}

// CONTAINS
func Test_NumericSliceContainsIntsReturnsTrue(t *testing.T) {
	numbers := NumericSlice[int]{0, 2}

	result := numbers.Contains(0)

	assert.True(t, result)
}

func Test_NumericSliceContainsIntsReturnFalse(t *testing.T) {
	numbers := NumericSlice[int]{0, 2}

	result := numbers.Contains(-1)

	assert.False(t, result)
}

// WHERE
func Test_NumericWhereIntsReturnsEven(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.Where(isEven)

	assert.Equal(t, result, NumericSlice[int]{0, 2, 4, 6, 8})
}

func Test_NumericWhereIntsReturnsOdd(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := numbers.Where(isOdd)

	assert.Equal(t, result, NumericSlice[int]{1, 3, 5, 7, 9})
}

// EXCEPT
func Test_NumericExceptIntsReturnsNotEven(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := numbers.Except(isEven)

	assert.Equal(t, NumericSlice[int]{1, 3, 5, 7, 9}, result)
}

func Test_NumericExceptIntsReturnsNotOdd(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := numbers.Except(isOdd)

	assert.Equal(t, NumericSlice[int]{0, 2, 4, 6, 8}, result)
}

// FIRST
func Test_NumericFirstIntsReturnsEven(t *testing.T) {
	numbers := NumericSlice[int]{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := numbers.First(isEven)

	assert.Nil(t, err)
	assert.Equal(t, result, 0)
}

func Test_NumericFirstIntsReturnsNotFound(t *testing.T) {
	numbers := NumericSlice[int]{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := numbers.First(isOdd)

	assert.NotNil(t, err)
	assert.Equal(t, err, golinq.NotFoundError{})
	assert.Equal(t, result, 0)
}

// LAST
func Test_NumericLastIntsReturnsEven(t *testing.T) {
	numbers := NumericSlice[int]{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := numbers.Last(isEven)

	assert.Nil(t, err)
	assert.Equal(t, result, 2)
}

func Test_NumericLastIntsReturnsNotFound(t *testing.T) {
	numbers := NumericSlice[int]{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := numbers.Last(isOdd)

	assert.NotNil(t, err)
	assert.Equal(t, err, golinq.NotFoundError{})
	assert.Equal(t, result, 0)
}

// TAKE
func Test_NumericTakeIntsReturnsFirst5(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Take(5)

	assert.Equal(t, NumericSlice[int]{0, 1, 2, 3, 4}, result)
}

func Test_NumericTakeIntsReturnsAllAvailable(t *testing.T) {
	numbers := NumericSlice[int]{0, 1}

	result := numbers.Take(5)

	assert.Equal(t, NumericSlice[int]{0, 1}, result)
}

func Test_NumericTakeIntsReturnsAllAvailableNil(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Take(5)

	var expected NumericSlice[int]
	assert.Equal(t, expected, result)
}

// SKIP
func Test_NumericSkipIntsReturnsExceptFirst5(t *testing.T) {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Skip(5)

	assert.Equal(t, Slice[int]{5, 6, 7, 8, 9}, result)
}

func Test_NumericSkipIntsReturnsAllAvailableExceptFirst5(t *testing.T) {
	numbers := NumericSlice[int]{0, 1}

	result := numbers.Skip(5)

	var expected NumericSlice[int]
	assert.Equal(t, expected, result)
}

// ORDER BY
func Test_NumericOrderByIntsReturnsAscending(t *testing.T) {
	numbers := NumericSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	lessThan := func(i, j int) bool {
		return numbers[i] < numbers[j]
	}

	result := numbers.OrderBy(lessThan)

	assert.Equal(t, NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}

// MAX
func Test_NumericMaxIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := numbers.Max()

	assert.Equal(t, 9, result)
}

func Test_NumericMaxIntsReturnsNil(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Max()

	assert.Equal(t, 0, result)
}

// MIN
func Test_NumericMinIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := numbers.Min()

	assert.Equal(t, 0, result)
}

func Test_NumericMinIntsReturnsLast(t *testing.T) {
	numbers := NumericSlice[int]{3, 7, 6, 9, 8, 0, 4, 2, 1, 5, -1}

	result := numbers.Min()

	assert.Equal(t, -1, result)
}

func Test_NumericMinIntsReturnsNil(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Min()

	assert.Equal(t, 0, result)
}

// SUM
func Test_NumericSumIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Sum()

	assert.Equal(t, 45, result)
}

func Test_NumericSumIntsReturnsNil(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Sum()

	assert.Equal(t, 0, result)
}

// AVERAGE
func Test_NumericAverageIntsReturns(t *testing.T) {
	numbers := NumericSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := numbers.Average()

	assert.Equal(t, 4.5, result)
}

func Test_NumericAverageIntsReturnsNil(t *testing.T) {
	numbers := NumericSlice[int]{}

	result := numbers.Average()

	assert.Equal(t, float64(0), result)
}
