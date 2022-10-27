package shadow

import (
	"fmt"
)

func Shadow() {
	fmt.Println("Shaddow")
	x := 5
	if x <= 5 {
		x := 7
		fmt.Println(x) // 7
	}
	fmt.Println(x) // 5
}

func Shadow2() {
	fmt.Println("Shaddow 2")
	x := 5
	if x <= 5 {
		x = 7
		fmt.Println(x) // 7
	}
	fmt.Println(x) // 7
}
