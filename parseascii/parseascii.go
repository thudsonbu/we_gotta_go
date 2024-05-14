package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	const ascii = "86 101 114 121 32 104 101 108 112 102 117 108 32 104 101 104 101"
	const base = 10
	const bitSize = 8
	slice := split(ascii, " ")
	ints := convert(slice, base, bitSize)
	str := convertToString(ints)
	fmt.Println(str)
}

// split splits a string into a slice of strings.
func split(s, sep string) []string {
	return strings.Split(s, sep)
}

// convert converts a slice of strings into a slice of integers.
func convert(slice []string, base, bitSize int) []int {
	ints := make([]int, len(slice))
	for i, s := range slice {
		n, err := strconv.ParseInt(s, base, bitSize)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = int(n)
	}
	return ints
}

// convertToString converts a slice of integers into a string.
func convertToString(ints []int) string {
	runes := make([]rune, len(ints))
	for i, n := range ints {
		runes[i] = rune(n)
	}
	return string(runes)
}
