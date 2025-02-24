//Problem: Create a struct named Car with fields Brand, Model, and Year.
// Write a function NewCar that takes these values as arguments and returns a Car instance.
// Then, create a PrintCarDetails function that prints the details of a car.

package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

func NewCar(brand, model string, year int) Car {
	return Car{Brand: brand, Model: model, Year: year}
}

func PrintCarDetails(car Car) {
	fmt.Printf("Brand: %s, Model: %s, Year: %d\n", car.Brand, car.Brand, car.Year)
}

func main() {
	car := NewCar("Tesla", "Model 3", 25)
	PrintCarDetails(car)
}
