package main

import "strings"

var justString string

// The negative consequence of the code snippet given in the task
// is that it creates a huge string of size 1 << 10,
// but only the first 100 characters are stored in the justString variable.
// The remaining part of the huge string remains in memory until the program terminates,
// even though it is no longer needed.
// We waste memory space!
// To fix this issue, it is better to avoid creating a
// huge string in the first place, and only create a string of the required size.
// One way to achieve this is to use the strings.Builder type,
// which provides a more efficient way to build strings dynamically
func someFunc() {
	var builder strings.Builder
	builder.Grow(100) // set the initial size of the string to 100 bytes
	builder.WriteString(createSmallString())
	justString = builder.String() // get the final string value
}

func createSmallString() string {
	// create and return a small string
	panic("Not yet implemented")
}

func main() {
	someFunc()
}
