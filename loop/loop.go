package loop

import (
	"fmt"
)

func Loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i) // 0 1 2 3 4
	}

	i := 0
	for i < 5 {
		fmt.Println(i) // 0 1 2 3 4
	}

	j := 0
	for {
		if j >= 5 {
			break
		}
		fmt.Println(j) // 0 1 2 3 4
		j++
	}

var arr = [5]int{1, 2, 3, 4, 5}

for i, v := range arr {
	fmt.Println(i, v) // 0 1 1 2 2 3 3 4 4 5
}

// use an underscore to ignore the index
for _, v := range arr {
	fmt.Println(v) // 1 2 3 4 5
}
}
