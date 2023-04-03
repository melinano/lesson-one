package main

import (
	"fmt"
	"time"
)

// sleep function takes a time.Duration argument
// inside the function, we block on the channel returned by time.After for the specified duration
// when the duration has elapsed, the channel receives a signal,
// and we can continue executing the rest of the program.
func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Start")
	sleep(2 * time.Second)
	fmt.Println("End")
}
