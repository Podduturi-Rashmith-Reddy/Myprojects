package main

import (
	"fmt"
)

// Vehicle interface
type Vehicle interface {
	Drive()
}

// Engine struct with a StartEngine method
type Engine struct{}

// StartEngine method for Engine
func (e Engine) StartEngine() {
	fmt.Println("Engine started...")
}

// Car struct embedding Engine
type Car struct {
	Engine
	Brand string
}

// Implement Drive() for Car
func (c Car) Drive() {
	c.StartEngine() // Call embedded method
	fmt.Printf("Driving the %s car...\n", c.Brand)
}

// Bike struct
type Bike struct {
	Brand string
}

// Implement Drive() for Bike
func (b Bike) Drive() {
	fmt.Printf("Riding the %s bike...\n", b.Brand)
}

func main() {
	// Create instances of Car and Bike
	myCar := Car{Engine: Engine{}, Brand: "Toyota"}
	myBike := Bike{Brand: "Yamaha"}

	// Vehicle interface slice
	vehicles := []Vehicle{myCar, myBike}

	// Iterate and drive each vehicle
	for _, v := range vehicles {
		v.Drive()
	}
}
