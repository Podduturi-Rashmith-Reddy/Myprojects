package main

import "fmt"

func main() {

	numbers := []int{10, 20, 30, 40, 50}

	numbers = append(numbers, 60)

	for i, num := range numbers {
		fmt.Println("Index:", i, "Value:", num)
	}

	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}
