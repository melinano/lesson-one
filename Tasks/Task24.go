package main

import (
	"fmt"
	"math"
)

// Point struct
type Point struct {
	x float64
	y float64
}

// constructor of Point struct
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// member function to calculate the distance to another point
func (p *Point) DistanceTo(other *Point) float64 {
	// using the pythagorean theorem to get the distance
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// creating to example points
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)
	// retrieving the distance
	distance := point1.DistanceTo(point2)
	// printing the result
	fmt.Printf("The distance between point1 and point2 is %.2f\n", distance)
}
