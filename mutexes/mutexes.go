package mutexes

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Mutexes() {
	// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an
	// unlocked mutex.
	var m sync.Mutex

	// Lock locks m. If the lock is already in use, the calling goroutine blocks
	// until the mutex is available.
	m.Lock()

	// Unlock unlocks m. It is a run-time error if m is not locked on entry to
	// Unlock.
	m.Unlock()

	// A Mutex can be locked more than once. It is an error if a mutex is unlocked
	// too many times.
	m.Lock()
	m.Lock()
	m.Unlock()
	m.Unlock()
}

// compare and swap
func CAS() {
	// CompareAndSwapInt32 executes the compare-and-swap operation for an int32
	// value.
	var i int32
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			v := i
			if atomic.CompareAndSwapInt32(&i, v, v+1) {
				break
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			v := i
			if atomic.CompareAndSwapInt32(&i, v, v+1) {
				break
			}
		}
	}()
	wg.Wait()
	fmt.Println(i)
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

type SafeCounterCAS struct {
	v map[string]*int32
}

// Inc increments the counter for the given key.
func (c *SafeCounterCAS) Inc(key string) {
	var oldValue, newValue int32
	for {
		oldValue = *c.v[key]
		newValue = oldValue + 1
		if atomic.CompareAndSwapInt32(c.v[key], oldValue, newValue) {
			break
		}
	}
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounterCAS) Value(key string) int {
	return int(*c.v[key])
}
