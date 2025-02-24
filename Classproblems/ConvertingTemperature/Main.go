package main

import "fmt"

type Temperature struct {
	Value float64
}

func (t *Temperature) ToFahrenheit() {
	t.Value = (t.Value * 9 / 5) + 32
}

func (t *Temperature) ToCelsius() {
	t.Value = (t.Value - 32) * 5 / 9
}

func main() {
	temp := Temperature{Value: 25}
	fmt.Println("Celsius:", temp.Value)

	temp.ToFahrenheit()
	fmt.Println("Converted to Fahrenheit:", temp.Value)

	temp.ToCelsius()
	fmt.Println("Converted back to Celsius:", temp.Value)
}
