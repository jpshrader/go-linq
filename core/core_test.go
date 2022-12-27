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

	assert.Equal(t, []int{0, 2, 4, 6, 8}, result)
}

func Test_WhereIntsReturnsOdd(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := Where(numbers, isOdd)

	assert.Equal(t, []int{1, 3, 5, 7, 9}, result)
}

// EXCEPT
func Test_ExceptIntsReturnsNotEven(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result := Except(numbers, isEven)

	assert.Equal(t, []int{1, 3, 5, 7, 9}, result)
}

func Test_ExceptIntsReturnsNotOdd(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result := Except(numbers, isOdd)

	assert.Equal(t, []int{0, 2, 4, 6, 8}, result)
}

// FIRST
func Test_FirstIntsReturnsEven(t *testing.T) {
	numbers := []int{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := First(numbers, isEven, -1)

	assert.Nil(t, err)
	assert.Equal(t, 0, result)
}

func Test_FirstIntsReturnsNotFound(t *testing.T) {
	numbers := []int{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := First(numbers, isOdd, -1)

	assert.NotNil(t, err)
	assert.Equal(t, err, errors.NotFoundError{})
	assert.Equal(t, -1, result)
}

// LAST
func Test_LastIntsReturnsEven(t *testing.T) {
	numbers := []int{0, 2}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	result, err := Last(numbers, isEven, -1)

	assert.Nil(t, err)
	assert.Equal(t, 2, result)
}

func Test_LastIntsReturnsNotFound(t *testing.T) {
	numbers := []int{0, 2}

	isOdd := func(x int) bool {
		return x%2 == 1
	}

	result, err := Last(numbers, isOdd, -1)

	assert.NotNil(t, err)
	assert.Equal(t, err, errors.NotFoundError{})
	assert.Equal(t, -1, result)
}

// TAKE
func Test_TakeIntsReturnsFirst5(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := Take(numbers, 5)

	assert.Equal(t, []int{0, 1, 2, 3, 4}, result)
}

func Test_TakeIntsReturnsAllAvailable(t *testing.T) {
	numbers := []int{0, 1}

	result := Take(numbers, 5)

	assert.Equal(t, []int{0, 1}, result)
}

func Test_TakeIntsReturnsAllAvailableNil(t *testing.T) {
	numbers := []int{}

	result := Take(numbers, 5)

	var expected []int
	assert.Equal(t, expected, result)
}

// SKIP
func Test_SkipIntsReturnsExceptFirst5(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := Skip(numbers, 5)

	assert.Equal(t, []int{5, 6, 7, 8, 9}, result)
}

func Test_SkipIntsReturnsAllAvailableExceptFirst5(t *testing.T) {
	numbers := []int{0, 1}

	result := Skip(numbers, 5)

	var expected []int
	assert.Equal(t, expected, result)
}

// SELECT
type person struct {
	name string
}

func personTransformer(x person) string {
	return x.name
}

func Test_SelectReturns(t *testing.T) {
	p1 := person{
		name: "Testing",
	}
	p2 := person{
		name: "test",
	}
	p3 := person{
		name: "gnitset",
	}

	persons := []person{p1, p2, p3}

	result := Select(persons, personTransformer)

	assert.Equal(t, []string{p1.name, p2.name, p3.name}, result)
}

func Test_SelectReturnsNil(t *testing.T) {
	persons := []person{}

	result := Select(persons, personTransformer)

	var expected []string
	assert.Equal(t, expected, result)
}

// DISTINCT
func Test_DistinctIntsReturns(t *testing.T) {
	numbers := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5}

	result := Distinct(numbers)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, result)
}

func Test_DistinctIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	result := Distinct(numbers)

	var expected []int
	assert.Equal(t, expected, result)
}

// DISTINCT OBJECT
func Test_DistinctStructIntsReturns(t *testing.T) {
	numbers := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5}

	transform := func(x int) int {
		return x
	}

	result := DistinctC(numbers, transform)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, result)
}

func Test_DistinctStructIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	transform := func(x int) int {
		return x
	}

	result := DistinctC(numbers, transform)

	var expected []int
	assert.Equal(t, expected, result)
}

func Test_DistinctStructReturns(t *testing.T) {
	p1 := person{
		name: "Testing",
	}
	p2 := person{
		name: "test",
	}
	p3 := person{
		name: "gnitset",
	}

	persons := []person{p1, p2, p1, p2, p3}

	result := DistinctC(persons, personTransformer)

	assert.Equal(t, []person{p1, p2, p3}, result)
}

func Test_DistinctStructReturnsNil(t *testing.T) {
	persons := []person{}

	result := DistinctC(persons, personTransformer)

	var expected []person
	assert.Equal(t, expected, result)
}

// ORDER BY
func Test_OrderByIntsReturnsAscending(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	lessThan := func(i, j int) bool {
		return numbers[i] < numbers[j]
	}

	result := OrderBy(numbers, lessThan)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}

func Test_OrderByIntsReturnsDescending(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	greaterThan := func(i, j int) bool {
		return numbers[i] > numbers[j]
	}

	result := OrderBy(numbers, greaterThan)

	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, result)
}

// ORDER BY ASCENDING
func Test_OrderByAscendingIntsReturns(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := OrderByAscending(numbers)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}

// ORDER BY DESCENDING
func Test_OrderByDescendingIntsReturns(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := OrderByDescending(numbers)

	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, result)
}

// MAX
func getNum(x int) int {
	return x
}

func Test_MaxIntsReturns(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := Max(numbers)

	assert.Equal(t, 9, result)
}

func Test_MaxIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	result := MaxC(numbers, getNum)

	assert.Equal(t, 0, result)
}

func Test_MaxCIntsReturns(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := MaxC(numbers, getNum)

	assert.Equal(t, 9, result)
}

func Test_MaxCIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	result := Max(numbers)

	assert.Equal(t, 0, result)
}

// MIN
func Test_MinIntsReturns(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := Min(numbers)

	assert.Equal(t, 0, result)
}

func Test_MinIntsReturnsLast(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5, -1}

	result := Min(numbers)

	assert.Equal(t, -1, result)
}

func Test_MinIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	result := Min(numbers)

	assert.Equal(t, 0, result)
}

func Test_MinCIntsReturns(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5}

	result := MinC(numbers, getNum)

	assert.Equal(t, 0, result)
}

func Test_MinCIntsReturnsLast(t *testing.T) {
	numbers := []int{3, 7, 6, 9, 8, 0, 4, 2, 1, 5, -1}

	result := MinC(numbers, getNum)

	assert.Equal(t, -1, result)
}

func Test_MinCIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	result := MinC(numbers, getNum)

	assert.Equal(t, 0, result)
}

// SUM
func Test_SumIntsReturns(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := Sum(numbers)

	assert.Equal(t, 45, result)
}

func Test_SumIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	result := Sum(numbers)

	assert.Equal(t, 0, result)
}

func Test_SumCIntsReturns(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := SumC(numbers, getNum)

	assert.Equal(t, 45, result)
}

func Test_SumCIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	result := SumC(numbers, getNum)

	assert.Equal(t, 0, result)
}

// AVERAGE
func Test_AverageIntsReturns(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := Average(numbers)

	assert.Equal(t, 4.5, result)
}

func Test_AverageIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	result := Average(numbers)

	assert.Equal(t, float64(0), result)
}

func Test_AverageCIntsReturns(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := AverageC(numbers, getNum)

	assert.Equal(t, 4.5, result)
}

func Test_AverageCIntsReturnsNil(t *testing.T) {
	numbers := []int{}

	result := AverageC(numbers, getNum)

	assert.Equal(t, float64(0), result)
}
