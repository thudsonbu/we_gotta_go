package types

import "testing"

func TestBirthday(t *testing.T) {
	tom := Person{
		Name: "Tom",
		Age:  23,
	}

	tom.Birthday()

	if tom.Age != 24 {
		t.Errorf("Expected age to be 24, got %d", tom.Age)
	}
}

func TestScore(t *testing.T) {
	s := Score{
		score: 0,
		name:  "Tom",
	}

	if s.increment() != 1 {
		t.Errorf("Expected score to be 1, got %d", s.score)
	}
}
