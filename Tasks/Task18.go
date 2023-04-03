package main

import (
	"fmt"
	"sync"
)

// Counter struct
type Counter struct {
	mu    sync.Mutex
	value int
}

// memeber-function of the struct,
// that increments its value
func (c *Counter) Increment() {
	// making sure that only one thread at a time can access
	// the Counter instance
	c.mu.Lock()
	defer c.mu.Unlock()
	// increment the value
	c.value++
}

// memeber-function of the struct,
// that returns its valuea as an integer
func (c *Counter) Value() int {
	// making sure that only one thread at a time can access
	// the Counter instance
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// create WaitGroup for synchronizing goroutines
	var wg sync.WaitGroup
	// create a Counter
	counter := Counter{}

	// create a thousand goroutines that run concurrently and
	// increment the counter
	for i := 0; i < 1000; i++ {
		// increment WaitGroup counter
		wg.Add(1)
		go func() {
			// invoking the increment-function
			counter.Increment()
			// decrement WaitGroup counter
			wg.Done()
		}()
	}

	// waiting for all goroutines to finish
	wg.Wait()
	// output the value of the counter
	fmt.Printf("Counter value: %d\n", counter.Value())
}
