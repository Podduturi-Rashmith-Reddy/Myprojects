package main

import "fmt"

func main() {

	var numbers [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Println("First Element:", numbers[0])

	for i, num := range numbers {
		fmt.Println("Index:", i, "Value:", num)
	}
}
