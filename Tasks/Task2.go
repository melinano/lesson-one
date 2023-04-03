package main

import (
	"fmt"
	"sync"
)

func main() {
	// instantiating slice
	nums := []int{2, 4, 6, 8, 10}

	// Create a Waitgroup to synchronize the goroutines
	var wg sync.WaitGroup

	// Iterate over the array and spawn a goroutine for each number
	for _, num := range nums {
		// adding goroutine function to Waitgroup
		// and therefore incrementing the Waitgroup counter by 1
		wg.Add(1)
		go func(n int) {
			// calculating the square value
			square := n * n
			// printing the result
			fmt.Printf("%d^2 = %d\n", n, square)
			// finishing goroutine and therefore decrement
			// the Waitgroup counter
			wg.Done()
		}(num)
	}

	// Wait for all the goroutines to finish
	wg.Wait()
}
