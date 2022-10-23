package main

import (
	"fmt"
	"hello/morestrings"
)

func main() {
	msg := morestrings.ReverseRunes("Hello, me.");
	fmt.Println(msg)
}
