package main

import (
	"fmt"
	"reflect"

	golinq "github.com/jpshrader/go-linq"
	"github.com/jpshrader/go-linq/slices"
)

func main() {
	numberExample()
	constraintExample()
	typeExample()
}

// NUMBERS EXAMPLE
func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 == 1
}

func numberExample() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	evenNums := golinq.Where(numbers, isEven)

	nums := slices.AsNumericSlice(numbers)
	oddNums := nums.Where(isOdd)

	linqFun := oddNums.
		Concat(evenNums).
		Skip(2).
		Reverse().
		Skip(2).
		Take(5).
		OrderByAscending()

	fmt.Println("================= NUMBERS EXAMPLE =================")
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("even numbers: %v\n", evenNums)
	fmt.Printf("odd numbers: %v\n", oddNums)
	fmt.Printf("linq fun: %v\n", linqFun)
}

// CONSTRAINTS EXAMPLE
func constraintExample() {
	numbers := []int{9, 6, 5, 4, 8, 2, 7, 3, 1, 0}
	strings := []string{"j", "f", "g", "c", "a", "h", "b", "e", "i", "d"}

	nums := slices.AsNumericSlice(numbers)
	strs := slices.AsOrderedSlice(strings)

	fmt.Println("================= CONSTRAINTS EXAMPLE =================")
	fmt.Printf("nums: %d\n", nums.OrderByAscending())
	fmt.Printf("sum nums: %d\n", nums.Sum())
	fmt.Printf("max num: %d\n", nums.Max())
	fmt.Printf("letters: %s\n", strs.OrderByAscending())
	item, _ := strs.OrderByDescending().First(nil)
	fmt.Printf("letter count: %s\n", item)
}

// TYPE EXAMPLE
func typeExample() {
	persons := []Person{
		{Name: "John", Age: 26},
		{Name: "Jack", Age: 23},
	}
	employees := []Employee{
		{Name: "Jack", Age: 26},
		{Name: "John", Age: 23},
	}

	above25 := func(p Person) bool {
		return p.Age >= 25
	}
	personsSlice := slices.AsSlice(persons)
	fmt.Println("persons under 25 years old: ", personsSlice.Except(above25))
	fmt.Println("persons above 25 years old: ", personsSlice.Where(above25))

	getPersonNames(persons)

	getNames(persons)
	people := getNamesInterfaceMethod(persons)

	getNames(employees)
	emp := getNamesInterfaceMethod(employees)

	fmt.Println("================= PEOPLE EXAMPLE =================")
	fmt.Printf("people: %v\n", people)
	fmt.Printf("employees: %v\n", emp)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) GetName() string {
	return p.Name
}

type Employee struct {
	Name string
	Age  int
}

func (e Employee) GetName() string {
	return e.Name
}

type PersonOrEmployee interface {
	Person | Employee
	GetName() string
}

func getPersonNames[T Person](people []T) []string {
	return golinq.Select(people, func(p T) string {
		return Person(p).Name
	})
}

func getNames[T Person | Employee](people []T) []string {
	return golinq.Select(people, func(p T) string {
		switch reflect.TypeOf(p) {
		case reflect.TypeOf(Person{}):
			return Person(p).Name
		case reflect.TypeOf(Employee{}):
			return Employee(p).Name
		default:
			return ""
		}
	})
}

func getNamesInterfaceMethod[T PersonOrEmployee](people []T) []string {
	return golinq.Select(people, func(p T) string {
		return p.GetName()
	})
}
