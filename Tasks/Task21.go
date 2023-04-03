package main

import "fmt"

// transport interface implementing a certain function
type transport interface {
	navigateToDestination()
}

// struct to simulate a requirement of "Car needs to go on water"
type Client struct {
}

// the client.startNavigation requires a parameter of transport type
func (t *Client) startNavigation(trans transport) {
	fmt.Println("client starting the navigation process")
	trans.navigateToDestination()
}

// Boat struct with certain functionalities
type Boat struct {
}

func (b *Boat) navigateToDestination() {
	fmt.Println("Boat is navigating to island")
}

// struct Car with different functionalities
// e.g. it is not able to drive on water
type Car struct {
}

func (c *Car) driveToDestination() {
	fmt.Println("Car is driving to destination")
}

// Adapter struct used by Car, so it
// can go to an island
type carAdapter struct {
	carTransport *Car
}

// it implements the navigateToDestination() function
// of the transport interface
func (c *carAdapter) navigateToDestination() {
	fmt.Println("Adapter modify Car to allow navigation.")
	// invoking the driveToDestination() function of the car instead
	c.carTransport.driveToDestination()
}

func main() {
	// creating a client and boat
	client := &Client{}
	boat := &Boat{}

	// start navigating on boat
	client.startNavigation(boat)

	// creating a car
	car := &Car{}
	// creating the carAdapter
	carAdapter := &carAdapter{
		carTransport: car,
	}

	// navigate on the car
	client.startNavigation(carAdapter)
}
