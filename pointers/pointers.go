package main

import "fmt"

// You can retrieve the pointer of a variable by prepending &, the address
// operator. You can retrieve the value of a pointer by prepending *
// (the indirection operator) to a pointer. This is called dereferencing.
func main() {
	x := 3

	xPointer := &x

	fmt.Println(xPointer) // 0x0023492384

	xVal := *xPointer

	fmt.Println(xVal) // 3
}
