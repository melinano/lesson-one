package main

import (
	"fmt"
	"sync"
)

func main() {
	// create a map with key and vale-type integer
	m := make(map[int]int)
	// declare a WaitGroup instance
	var wg sync.WaitGroup
	// declare a Mutex instance
	var mu sync.Mutex

	// Spawn 10 goroutines to write to the map concurrently
	for i := 0; i < 10; i++ {
		// increment WainGroup counter
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			// Acquire a lock before writing to the map
			// so only one goroutine can access the map at a time
			mu.Lock()
			// writing n as key and the square of n to the map as value
			m[n] = n * n
			// release the mutex lock
			mu.Unlock()
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the contents of the map
	fmt.Println(m)
}
