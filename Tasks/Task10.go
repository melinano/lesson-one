package main

import (
	"fmt"
	"math"
)

func main() {
	// slice with temperature fluctuations
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// empty dictionary where
	// key: temperature range
	// value: list of temperatures that fall in that range
	tempRanges := make(map[int][]float64)

	// iterating over all temperature values
	for _, temp := range temps {
		// dividing by 10 and rounding down to the nearest integer
		rangeKey := int(math.Floor(temp/10) * 10)
		// if this integer value already exists in the dictionary
		// we append the temperature to the list of temperatures for that range
		// if not - we create a new list
		if _, ok := tempRanges[rangeKey]; ok {
			tempRanges[rangeKey] = append(tempRanges[rangeKey], temp)
		} else {
			tempRanges[rangeKey] = []float64{temp}
		}
	}

	// Output the dictionary
	for rangeKey, temps := range tempRanges {
		fmt.Printf("%d:{%v}\n", rangeKey, temps)
	}
}
