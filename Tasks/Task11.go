package main

import "fmt"

func main() {
	// two disordered sets
	set1 := []int{1, 4, 5, 8, 13, 9, 43, 3}
	set2 := []int{1, 5, 2, 9, 55, 77, 42, 8}

	// using a map
	set1Map := make(map[int]bool)
	// iterate over the elements of the first set
	// and add them to the map
	for _, val := range set1 {
		set1Map[val] = true
	}

	// empty slice for the intersection
	intersection := []int{}

	// iterate over the second set and check if
	// the value exists in the map
	// if so, add them to the intersection slice
	for _, val := range set2 {
		if set1Map[val] {
			intersection = append(intersection, val)
		}
	}
	fmt.Println(intersection)
}
