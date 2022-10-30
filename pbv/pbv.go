package pbv

import "errors"

type numMap map[int]int

// Althought functions are pass by value, the values of maps are memory pointers
// so you can still modify them in functions
func ReZeroMap(m numMap) (numMap, error) {
	if len(m) == 0 {
		return m, errors.New("Cannot write to an empty map")
	}

	m[0] = 0

	return m, nil
}
