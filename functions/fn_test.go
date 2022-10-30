package fn

import (
	"errors"
	"testing"
)

func TestRunFunc(t *testing.T) {
	var goodFn = func(a int) (int, error) {
		return a, nil
	}

	var badFn = func(a int) (int, error) {
		return a, errors.New("Off by One")
	}

	type tStruct struct {
		fn   runnableFunc
		in   int
		want int
		err  error
	}

	cases := []tStruct{
		{fn: goodFn, in: 1, want: 1, err: nil},
		{fn: badFn, in: 1, want: 0, err: errors.New("Off by One")},
	}

	for _, c := range cases {
		got, _ := RunFunc(c.fn, c.in)

		if got != c.want {
			t.Errorf("runFunc(%q) output %q, expected %q", c.in, got, c.want)
		}
	}
}
