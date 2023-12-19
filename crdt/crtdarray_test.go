package crdt

import (
	"fmt"
	"sync"
	"testing"
)

func TestCRDTArray(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Create two CRDTArrays.
	a := NewCRDTArray("client-1")
	b := NewCRDTArray("client-2")

	// Add some elements to each.
	go func() {
		defer wg.Done()
		a.Add("a")
		a.Add("b")
		a.Add("c")
	}()

	go func() {
		defer wg.Done()
		b.Add("d")
		b.Add("e")
		b.Add("f")
	}()

	// Wait for the goroutines to finish.
	wg.Wait()

	// Merge them together.
	a.Merge(b)
	b.Merge(a)

	output := "Array A: "
	for _, element := range a.data {
		output += fmt.Sprintf("%s ", element.Value)
	}
	output += "\nArray B: "
	for _, element := range b.data {
		output += fmt.Sprintf("%s ", element.Value)
	}

	if output != "Array A: a b c d e f \nArray B: d e f a b c " {
		t.Errorf("expected %s, got %s", "Array A: a b c d e f \nArray B: d e f a b c ", output)
	}
}
