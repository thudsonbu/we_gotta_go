package pbv

import "testing"

func TestReZeroMap(t *testing.T) {
	goodMap := numMap{
		0: 1,
		1: 0,
	}

	newMap, err := ReZeroMap(goodMap)

	newVal, ok := newMap[0]

	if !ok {
		t.Error("ReZeroMap Removed zero key")
	}

	if newVal != 0 {
		t.Error("ReZeroMap did not set zero key value to zero")
	}

	if err != nil {
		t.Errorf("ReZeroMap returned an error with a valid map, error: %q", err)
	}

	var badMap map[int]int

	newMap, err = ReZeroMap(badMap)

	if err == nil {
		t.Error("ReZeroMap did not return an error with an empty map")
	}

	newVal, ok = newMap[0]

	if ok {
		t.Error("ReZeroMap set the zero key when it did not exist")
	}
}

func TestZeroInt(t *testing.T) {
	var x int = 3

	y := ReZeroInt(x)

	if x != 3 {
		t.Error("ReZeroInt modified the input value")
	}

	if y != 0 {
		t.Errorf("ReZeroInt did not return 0. It returned:")
		t.Error(y)
	}
}

func TestReZeroSlice(t *testing.T) {
	s := []int{1, 2, 3}

	zeroed := ReZeroSlice(s)

	for i, v := range zeroed {
		if v != 0 {
			t.Error("ReZeroSlice did not set zero at index:")
			t.Error(i)
		}
	}

	for i, v := range []int{0, 0, 0} {
		if s[i] != v {
			t.Error("ReZeroSlice did not modify original slice at index:")
			t.Error(i)
		}
	}
}
