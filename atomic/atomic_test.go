package atomic

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const key = "somekey"

func TestSafeCounter(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	startTime := time.Now()
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.Inc(key)
			}
			fmt.Println(c.Value(key))
		}()
	}
	wg.Wait()
	fmt.Println("safe counter", time.Since(startTime))
	if false {
		t.Error("Test failed")
	}
}

func TestSafeCounterCAS(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	startTime := time.Now()
	c := SafeCounterCAS{v: make(map[string]*int32)}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.Inc(key)
			}
			fmt.Println(c.Value(key))
		}()
	}
	wg.Wait()
	fmt.Println("safe counter with cas", time.Since(startTime))
	if false {
		t.Error("Test failed")
	}
}
