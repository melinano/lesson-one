package main

import "fmt"

func deleteElement(slice []int, i int) []int {
	// Remove the i-th element from the slice
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2                          // index of the element to delete
	slice = deleteElement(slice, i) // invoke the function to delete the i-th element
	fmt.Println(slice)
}
