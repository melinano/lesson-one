package main

import "fmt"

func main() {
	// creating two variables with different values
	a, b := 10, 20
	// swapping them with the help of
	// the "short assignment statement"
	a, b = b, a
	// printing the result
	fmt.Println(a, b)
}
