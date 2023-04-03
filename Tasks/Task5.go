package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel of integers
	ch := make(chan int)

	// Start a goroutine to send values to the channel
	go func() {
		for i := 1; ; i++ {
			ch <- i                            // sending i-value to the channel
			time.Sleep(100 * time.Millisecond) // Pause for 100 milliseconds
		}
	}()

	// Read values from the channel until the timeout occurs
	duration := 5 * time.Second // Set the duration to 5 seconds
	timeout := time.After(duration)
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				// Channel is closed, exit the loop
				fmt.Println("Channel closed, exiting...")
				return
			}
			fmt.Println("Received value:", val)
		case <-timeout:
			// Timeout, exit the loop
			fmt.Println("Timeout reached, exiting...")
			close(ch)
			return
		}
	}
}
