package main

import "fmt"

func reverseString(s string) string {
	// convert the string to a rune
	r := []rune(s)
	// iterating over the rune from opposite sides and swapping them
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	// return the flipped rune as string
	return string(r)
}

func main() {
	// example string
	str := "главрыба"

	// print out the flipped string
	fmt.Println(reverseString(str))
}
