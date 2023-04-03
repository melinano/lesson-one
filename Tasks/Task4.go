package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Create a channel for data of type string
	dataCh := make(chan string)

	// Start a constant data writer
	go func() {
		for i := 0; ; i++ {
			dataCh <- fmt.Sprintf("Data #%d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Creating a separate signal channel for the (CTRL+C) interrupt
	// so that it does not block the main channel
	sigCh := make(chan os.Signal, 1)
	// Wait for CTRL+C (syscall.SIGTERM) signal to terminate the program
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Select the number of workers
	numWorkers := 3

	// Create a waitgroup to synchronize the workers
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start the worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			// postponing the decrement of
			// waitgroup counter
			defer wg.Done()
			for data := range dataCh {
				fmt.Printf("Worker #%d received data: %s\n", id, data)
			}
			fmt.Printf("Worker #%d terminated\n", id)
		}(i)
	}

	// Wait for the Ctrl+C signal or for all workers to terminate
	select {
	case <-sigCh:
		// Interrupt signal (Ctrl+C) received
		fmt.Println("Program terminated by user")
	case <-waitWorkers(&wg):
		// All workers have terminated on their own
		fmt.Println("All workers have terminated")
	}

	// Close the data channel to signal to the workers to terminate
	close(dataCh)

	// Wait for all workers to terminate
	wg.Wait()

	fmt.Println("Program terminated")
}

// waitWorkers waits for all the workers to terminate and returns a channel (doneCh)
// that is closed when all workers have terminated.
func waitWorkers(wg *sync.WaitGroup) chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneCh)
	}()
	return doneCh
}
