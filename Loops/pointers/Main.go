package main

import "fmt"

func main() {
	num := 42
	ptr := &num

	fmt.Println("Original Value:", num)
	fmt.Println("Pointer Address:", ptr)
	fmt.Println("Value from Pointer:", *ptr)

	*ptr = 100
	fmt.Println("Modified Value:", num)
}
