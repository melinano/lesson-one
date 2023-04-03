package main

import (
	"fmt"
)

// Parent struct
type Human struct {
	// declaring struct fields
	Name   string
	Age    int
	Sex    string
	Height float32
	Weight int
}

// Member function of parent struct
func (human Human) ReturnFieldsAsString() string {
	return fmt.Sprintf(
		"My name is %s!\nI am %d years old, %s, %f tall and weight %dkg", human.Name, human.Age, human.Sex, human.Height, human.Weight)
}

// Child struct
type Action struct {
	// anonymous field of the parent struct
	Human
}

func main() {
	// Creating an instance of the child struct
	action := Action{
		// We can use the constructor of the parent struct here
		Human{
			Name:   "Peter",
			Age:    23,
			Sex:    "male",
			Height: 1.78,
			Weight: 100,
		},
	}
	// This gives the opportunity to invoke
	// member functions of the parent struct
	fmt.Println(action.ReturnFieldsAsString())
}
