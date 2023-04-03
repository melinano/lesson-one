package main

import "fmt"

func main() {
	// sequence of strings
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	// creating an empty map
	// key: string value
	// value: empty struct
	set := make(map[string]struct{})

	// assign the values from the string sequence to the map
	for _, s := range sequence {
		set[s] = struct{}{}
	}

	fmt.Println(set) // Print the result
}
