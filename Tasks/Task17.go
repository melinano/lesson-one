package main

import (
	"fmt"
	"sort"
)

func main() {
	// create a sorted slice of integers
	slice := []int{1, 3, 5, 7, 9, 11}

	// perform binary search on the slice
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= 5 // search for 5
	})

	// print the index of the found element
	if index < len(slice) && slice[index] == 5 {
		fmt.Printf("Found 5 at index %d\n", index)
	} else {
		fmt.Println("5 not found in slice")
	}
}
