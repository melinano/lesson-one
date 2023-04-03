package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 1st option:
// closing the channel and
// returning from the function
func main1() {
	ch := make(chan string, 6)
	go func() {
		for {
			v, ok := <-ch
			// if channel is closed the goroutine terminates
			if !ok {
				fmt.Println("Finish")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- "hello"
	ch <- "world"
	close(ch)
	time.Sleep(time.Second)
}

// 2nd option:
// periodic polling channel
func main2() {
	// creating channel with strings
	ch := make(chan string, 6)
	// creating variable of type channel
	// to be used as a semaphore to handle
	// the closing of the goroutine
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch <- "foo":
			// if signal comes from done channel
			// close the ch channel
			case <-done:
				// closing channel
				close(ch)
				// returning the function
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		// sending empty struct to signal the done channel
		done <- struct{}{}
	}()

	// printing the received messages from channel to console
	for i := range ch {
		fmt.Println("Received: ", i)
	}

	fmt.Println("Finish")
}

// 3rd option:
// using context

func main3() {
	// instantiating channel
	ch := make(chan struct{})
	// create a context
	ctx, cancel := context.WithCancel(context.Background())

	// start a goroutine with context
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if the context cease to exist
				// paste an empty struct inti the channel
				ch <- struct{}{}
				return
			default: // else print something
				fmt.Println("foo...")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	// concurrent function which closes the context
	// after 3 seconds
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("Finish")
}

// 4th option:
// using a waitgroup
func main() {
	// creating a WaitGroup
	var wg sync.WaitGroup
	// incrementing WaitGroup counter
	wg.Add(1)
	// goroutine function
	go func() {
		// decrement WaitGroup counter
		wg.Done()
	}()

	// wait for WaitGroup counter is zero
	wg.Wait()
}
