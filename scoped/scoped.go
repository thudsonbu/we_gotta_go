package scoped

import (
	"fmt"
)

func Scoped() {
	const x = 5

	if y, z := 9, 3; x == 5 { // y and z are scoped to the if statement
		fmt.Println(x, y, z) // 5 9 3
	}
}
