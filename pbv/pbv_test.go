package pbv

import (
	"testing"
)

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
		t.Error("ReZeroMap did not set zero key valye to zero")
	}

	if err != nil {
		t.Errorf("ReZeroMap returned an error with a valid map, error: %q", err )
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
