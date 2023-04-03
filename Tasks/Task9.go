package main

import "fmt"

func main() {
	// slice with example numbers
	nums := []int{1, 2, 3, 4, 5}

	// first channel
	ch1 := make(chan int)
	// second channel
	ch2 := make(chan int)

	// First channel writes numbers from slice
	// in a goroutine function
	go func() {
		defer close(ch1)
		// send each number from slice into the first channel
		for _, num := range nums {
			ch1 <- num
		}
	}()

	// Channel 2 multiplies the numbers by 2 and writes the results
	go func() {
		// postpone the closing of the channel till the end of function
		defer close(ch2)
		// multiply each number by 2 and send into the second channel
		for num := range ch1 {
			ch2 <- num * 2
		}
	}()

	// Read and output results from Channel 2
	for num := range ch2 {
		fmt.Println(num)
	}
}
