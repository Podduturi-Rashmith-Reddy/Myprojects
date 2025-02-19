package main

import (
	operations "Myprojects/Operations"
	"fmt"
)

func main() {
	a, b := 20, 10
	fmt.Println("Addition", operations.Add(a, b))
	fmt.Println("Subtraction", operations.Sub(a, b))
	fmt.Println("Multiplication", operations.Mul(a, b))
	result, err := operations.Div(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}
}
