package bitshifting

import "testing"

func TestMultiplyByTwo(t *testing.T) {
	if MultiplyByTwo(4) != 8 {
		t.Error("Expected 8")
	}
}

func TestDivideByTwo(t *testing.T) {
	if DivideByTwo(4) != 2 {
		t.Error("Expected 2")
	}
}

func TestToPower(t *testing.T) {
	if ToPower(4, 2) != 16 {
		t.Error("Expected 16")
	}

	if ToPower(4, 3) != 32 {
		t.Error("Expected 32")
	}
}

func TestToRoot(t *testing.T) {
	if ToRoot(4, 2) != 1 {
		t.Error("Expected 1")
	}

	if ToRoot(27, 3) != 3 {
		t.Error("Expected 3")
	}
}

func TestIsEven(t *testing.T) {
	if !IsEven(4) {
		t.Error("Expected true")
	}

	if IsEven(5) {
		t.Error("Expected false")
	}
}

func TestIsOdd(t *testing.T) {
	if !IsOdd(5) {
		t.Error("Expected true")
	}

	if IsOdd(4) {
		t.Error("Expected false")
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	if !IsPowerOfTwo(4) {
		t.Error("Expected true")
	}

	if IsPowerOfTwo(5) {
		t.Error("Expected false")
	}

	if IsPowerOfTwo(0) {
		t.Error("Expected false")
	}
}

func TestSwap(t *testing.T) {
	x, y := Swap(4, 5)
	if x != 5 || y != 4 {
		t.Error("Expected 5, 4")
	}
}
