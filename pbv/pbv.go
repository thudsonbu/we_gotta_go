package pbv

import "errors"

type numMap map[int]int

// Function in go are pass by value so even though we set x to 0 in this func
// the outer reference will not be changed
func ReZeroInt(x int) int {
	x = 0

	return x
}

// Althought functions are pass by value, the values of maps are memory pointers
// so you can still modify them in functions.
func ReZeroMap(m numMap) (numMap, error) {
	if len(m) == 0 {
		return m, errors.New("Cannot write to an empty map")
	}

	m[0] = 0

	return m, nil
}

// Slices are also implemented as pointers, so they can also be mutated within
// functions. However, their length cannot be changed.
func ReZeroSlice(s []int) []int {
	for i, _ := range s {
		s[i] = 0
	}

	return s
}
