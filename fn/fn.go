package fn

import "fmt"

func add(x int, y int) int {
	return x + y
}

// Simulate named and optional parameters with a struct
type Name struct {
	firstName string
	lastName string
}

func fullName(name Name) string {
	return name.firstName + " " + name.lastName
}

// Use variadic parameters when you don't know how many parameters you'll need
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	arr := [3]int{1, 2, 3}

	// Pass array to sum
	fmt.Println(sum(arr[:]...)) // 6
}

// Multiple return values are supported
func divAndRemainder(x int, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, fmt.Errorf("Cannot divide by zero")
	}
	return x / y, x % y, nil
}
