package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius
}

type Rectangle struct {
	Height, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func main() {
	var s Shape

	s = Circle{Radius: 5}
	fmt.Println("Circle Area = ", s.Area())

	s = Rectangle{Height: 4, Width: 5}
	fmt.Println("Rectangle Area = ", s.Area())
}
