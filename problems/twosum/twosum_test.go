package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, c := range cases {
		got := TwoSum(c.in, c.target)

		forward := got[0] != c.want[0] || got[1] != c.want[1]
		backward := got[0] != c.want[1] || got[1] != c.want[0]

		if !forward && !backward {
			t.Errorf("TwoSum(%q) == %q, expect %q", c.in, got, c.want)
		}
	}
}
