package main

import (
	"fmt"
	"strings"
)

func flipWords(s string) string {
	// splitting the string at the whitespaces into
	// a slice of strings
	words := strings.Fields(s)
	// iterating over the slice from opposite sides and swapping them
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// joining the separate string-words of
	// the slice to one string with whitespace between
	return strings.Join(words, " ")
}

func main() {
	// example string
	str := "snow dog sun"
	// printing out the result
	fmt.Println(flipWords(str))
}
