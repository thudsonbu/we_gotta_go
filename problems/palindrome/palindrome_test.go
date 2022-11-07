package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in  int
		exp bool
	}{
		{0, true},
		{1, true},
		{121, true},
		{10, false},
		{-121, false},
		{123, false},
		{1001, true},
		{10101, true},
		{32323, true},
	}

	for _, c := range cases {
		got := IsPalindrome(c.in)

		if got != c.exp {
			t.Errorf("IsPalindrome(%d) == %t, expected %t", c.in, got, c.exp)
		}
	}
}
