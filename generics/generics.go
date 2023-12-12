package generics

import (
	"fmt"
	"os"
	"reflect"
)

// PrintSlice is a generic function that works with a slice of any type.
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func printSlices() {
	PrintSlice([]int{1, 2, 3})          // works with slice of int
	PrintSlice([]string{"a", "b", "c"}) // works with slice of string
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Stringer interface {
	String() string
}

func checkWriter() {
	var w Writer = os.Stdout

	if s, ok := w.(Stringer); ok {
		fmt.Println(s.String())
	} else {
		fmt.Println("not a Stringer")
	}
}

func checkType[T int | string | bool](i T) string {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	default:
		return "unknown"
	}
}
