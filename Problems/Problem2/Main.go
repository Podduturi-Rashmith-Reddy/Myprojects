package main

import "fmt"

const OvenTime = 40

func RemainingOvenTime(actualMinustesInOven int) int {
	return OvenTime - actualMinustesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinustesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinustesInOven
}

func main() {
	fmt.Println("Preparation of Lasanga:")
	fmt.Println("Remaining Oven Time:", RemainingOvenTime(30))
	fmt.Println("Preparation Time:", PreparationTime(2))
	fmt.Println("Elapsed Time:", ElapsedTime(3, 20))
}
