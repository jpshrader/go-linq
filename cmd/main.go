package main

import (
	"fmt"
	"reflect"

	golinq "github.com/jpshrader/go-linq"
	"github.com/jpshrader/go-linq/receiver"
)

func main() {
	numberExample(true)
	constraintExample(true)
	typeExample(true)
}

// NUMBERS EXAMPLE
func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 == 1
}

func numberExample(print bool) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	evenNums := golinq.Where(numbers, isEven)

	nums := receiver.AsNumericSlice(numbers)
	oddNums := nums.Where(isOdd)

	linqFun := oddNums.
		Concat(evenNums).
		Skip(2).
		Reverse().
		Skip(2).
		Take(5).
		OrderByAscending()

	if print {
		fmt.Println("================= NUMBERS EXAMPLE =================")
		fmt.Printf("numbers: %v\n", numbers)
		fmt.Printf("even numbers: %v\n", evenNums)
		fmt.Printf("odd numbers: %v\n", oddNums)
		fmt.Printf("linq fun: %v\n", linqFun)
	}
}

// CONSTRAINTS EXAMPLE
func constraintExample(print bool) {
	numbers := []int{9, 6, 5, 4, 8, 2, 7, 3, 1, 0}
	strings := []string{"j", "f", "g", "c", "a", "h", "b", "e", "i", "d"}

	nums := receiver.AsNumericSlice(numbers)
	strs := receiver.AsOrderedSlice(strings)

	if print {
		fmt.Println("================= CONSTRAINTS EXAMPLE =================")
		fmt.Printf("nums: %d\n", nums.OrderByAscending())
		fmt.Printf("sum nums: %d\n", nums.Sum())
		fmt.Printf("max num: %d\n", nums.Max())
		fmt.Printf("letters: %s\n", strs.OrderByAscending())
		item, _ := strs.OrderByDescending().First(nil)
		fmt.Printf("letter count: %s\n", item)
	}
}

// TYPE EXAMPLE
func typeExample(print bool) {
	persons := []Person{
		{Name: "John", Age: 26},
	}
	employees := []Employee{
		{Name: "Jack", Age: 26},
	}

	getPersonNames(persons)

	getNames(persons)
	getNamesInterface(persons)
	people := getNamesInterfaceMethod(persons)

	getNames(employees)
	getNamesInterface(employees)
	emp := getNamesInterfaceMethod(employees)

	if print {
		fmt.Println("================= PEOPLE EXAMPLE =================")
		fmt.Printf("people: %v\n", people)
		fmt.Printf("employees: %v\n", emp)
	}
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

func getNamesInterface[T PersonOrEmployee](people []T) []string {
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
