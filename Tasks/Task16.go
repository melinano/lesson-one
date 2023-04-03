package main

import (
	"fmt"
	"sort"
)

// using the built-in sort package to quicksort a slice
func main() {
	// slice of numbers
	numbers := []int{5, 2, 4, 3, 1}
	// sorting the numbers
	sort.Ints(numbers)
	// printing them out
	fmt.Println(numbers)
}
