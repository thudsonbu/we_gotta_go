package numpalindrom

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
		{-121, true},
		{123, false},
		{-123, false},
	}

	for _, c := range cases {
		got := IsPalindrome(c.in)

		if got != c.exp {
			t.Errorf("IsPalindrome(%d) == %t, expected %t", c.in, got, c.exp)
		}
	}
}
