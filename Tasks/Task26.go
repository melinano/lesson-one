package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	// convert string to lowercase
	s = strings.ToLower(s)

	// create a map to store the frequency of each character
	charFreq := make(map[rune]int)

	// iterate through each character in the string
	for _, char := range s {
		// increment the frequency of the character in the map
		charFreq[char]++

		// if the frequency of the character is greater than 1, then the string is not unique
		if charFreq[char] > 1 {
			return false
		}
	}

	// if all characters have unique frequency, then return true
	return true
}

func main() {
	// test the function with different input strings
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}
