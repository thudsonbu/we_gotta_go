package shadow

import (
	"fmt"
)

// Variables modifications are block scoped when using := so x is only 7 in
// the below example for the length of the if block. This behavior is known as
// shadowing.
func Shadow() {
	fmt.Println("Shaddow")
	x := 5
	if x <= 5 {
		x := 7
		fmt.Println(x) // 7
	}
	fmt.Println(x) // 5
}

func NoShadow() {
	fmt.Println("Shaddow 2")
	x := 5
	if x <= 5 {
		x = 7
		fmt.Println(x) // 7
	}
	fmt.Println(x) // 7
}
