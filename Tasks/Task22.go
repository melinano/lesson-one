package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Initialize the two numeric variables with big.Int

	// shift the binary representation of the number 2 20 bits to the left
	a := new(big.Int).SetUint64(2 << 20)
	// shift the binary representation of the number 3 20 bits to the left
	b := new(big.Int).SetUint64(3 << 20)

	// using the bigInt functions to perform
	// the required operations
	// Multiply a and b
	mul := new(big.Int).Mul(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, mul)

	// Divide a by b
	div := new(big.Int).Div(a, b)
	fmt.Printf("%d / %d = %d\n", a, b, div)

	// Add a and b
	add := new(big.Int).Add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, add)

	// Subtract b from a
	sub := new(big.Int).Sub(a, b)
	fmt.Printf("%d - %d = %d\n", a, b, sub)

}
