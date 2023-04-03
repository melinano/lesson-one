package main

import (
	"fmt"
	"sync"
)

func main() {
	// instantiating slice
	nums := []int{2, 4, 6, 8, 10}

	// Create a waitgroup to synchronize the goroutines
	var wg sync.WaitGroup

	// Create a channel to receive the results from the goroutines
	resultCh := make(chan int)

	// Iterate over the array and spawn a goroutine for each number
	// and calculate their square value
	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			square := n * n
			// sending the square value to channel
			resultCh <- square
			wg.Done()
		}(num)
	}

	// Wait for all the goroutines to finish
	go func() {
		wg.Wait()
		// close the channel
		close(resultCh)
	}()

	// Sum up the results from the channel
	sum := 0
	for square := range resultCh {
		sum += square
	}

	// Output the result
	fmt.Printf("Sum of squares: %d\n", sum)
}
