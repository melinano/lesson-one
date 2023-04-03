package main

import "fmt"

func main() {
	// setting i-th bit to 1
	var x int64 = 123 // example variable
	i := 7            // example bit position

	oneMask := int64(1) << i // shift 1 to the left by i bits to create the mask
	x |= oneMask             // set the i-th bit of x to 1 using bitwise OR operator

	// setting j-th bit to 0
	var y int64 = 123 // example variable
	j := 3            // example bit position

	zeroMask := int64(^(1 << j)) // shift 1 to the left by j bits, invert the bits, and create the mask
	y &= zeroMask                // set the j-th bit of x to 0 using bitwise AND

	fmt.Printf("x with %d-th bit set to 1: %d\nand y with %d-th bit set to 0: %d", i, x, j, y)
}
