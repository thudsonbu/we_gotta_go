package crdt

import (
	"sync"
	"testing"
)

func TestCRDTKV(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Create two CRDTArrays.
	a := NewLWWKeyValueStore()
	b := NewLWWKeyValueStore()

	// Add some elements to each.
	go func() {
		defer wg.Done()
		a.Set("a", "a")
		a.Set("b", "b")
		a.Set("c", "c")
	}()

	go func() {
		defer wg.Done()
		b.Set("d", "d")
		b.Set("e", "e")
		b.Set("f", "f")
	}()

	// Wait for the goroutines to finish.
	wg.Wait()

	// Merge them together.
	a.Merge(b)
	b.Merge(a)

	letters := []string{"a", "b", "c", "d", "e", "f"}
	for _, letter := range letters {
		value, exists := a.Get(letter)
		if !exists {
			t.Errorf("expected %s to exist", letter)
		}
		if value != letter {
			t.Errorf("expected %s, got %s", letter, value)
		}
	}
}
