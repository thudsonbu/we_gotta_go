package concurrency

import (
	"fmt"
	"sync"
)

func WaitGroup() {
	// A WaitGroup is used to wait for a collection of goroutines to finish.
	var wg sync.WaitGroup

	// Add adds delta, which may be negative, to the WaitGroup counter.
	// If the counter becomes zero, all goroutines blocked on Wait are released.
	// If the counter goes negative, Add panics.
	wg.Add(2)

	// A goroutine is a lightweight thread managed by the Go runtime.
	go func() {
		// Decrement the counter when the goroutine completes.
		defer wg.Done()

		fmt.Println("Hello")
	}()

	go func() {
		// Decrement the counter when the goroutine completes.
		defer wg.Done()

		fmt.Println("World")
	}()

	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()
}

func Channels() {
	// A channel is a typed conduit through which you can send and receive values
	// with the channel operator, <-.
	// ch := make(chan int)
	ch := make(chan int, 2)

	// Send a value into a channel using the channel <- syntax.
	ch <- 1
	ch <- 2

	// Receive a value from a channel using the <-channel syntax.
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func BufferedChannels() {
	// A buffered channel is a channel that can hold more than one value.
	// The make function takes a second argument that specifies the buffer length.
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	// The buffered channel will block when the buffer is full.
	// ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func BlockingChannels() {
	// A blocking channel is a channel that will block until a value is received.
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)
}

// Select statements can solve the problem of blocking channels. The main
// goroutine is waiting for the inner goroutine to read the value it wrote to
// ch2 while the inner go routine is waiting on the main goroutine to read the
// value is wrote to ch1. This is a deadlock. The select statement will randomly
// choose a case to execute thus not blocking.
func Select() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()

	v := 2
	var v2 int

	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}

	fmt.Println(v, v2)
}
