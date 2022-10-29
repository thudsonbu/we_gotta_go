package fn

import "fmt"

func Add(x int, y int) int {
	return x + y
}

// Simulate named and optional parameters with a struct
type Name struct {
	firstName string
	lastName  string
}

func FullName(name Name) string {
	return name.firstName + " " + name.lastName
}

// Use variadic parameters when you don't know how many parameters you'll need
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func Main() {
	arr := [3]int{1, 2, 3}

	// Pass array to sum
	fmt.Println(Sum(arr[:]...)) // 6
}

// Multiple return values are supported
func DivAndRemainder(x int, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, fmt.Errorf("Cannot divide by zero")
	}
	return x / y, x % y, nil
}

// functions can have types
type runnableFunc func(a int) (int, error)

func RunFunc(fn runnableFunc, arg int) (int, error) {
	out, err := fn(arg)

	if err != nil {
		fmt.Println(err)

		return 0, err
	}

	return out, nil
}
