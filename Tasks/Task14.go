package main

import (
	"fmt"
	"reflect"
)

func main() {
	// creating variables of different types
	var i interface{} = 42
	checkType(i)

	var s interface{} = "hello"
	checkType(s)

	var b interface{} = true
	checkType(b)

	var ch interface{} = make(chan int)
	checkType(ch)
}

// takes a parameter of type interface{}
// interface{} is a special type that can hold any values
func checkType(i interface{}) {
	// retrieving the type of i
	// and printing the right type
	// with the help of a switch statement
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		fmt.Println("Type: int")
	case reflect.String:
		fmt.Println("Type: string")
	case reflect.Bool:
		fmt.Println("Type: bool")
	case reflect.Chan:
		fmt.Println("Type: channel")
	default:
		fmt.Println("Unknown type")
	}
}
