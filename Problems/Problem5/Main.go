package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * (successRate / 100.0) / 60)
}

// CalculateCost calculates the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10
	remainingCars := carsCount % 10

	return uint(groupsOfTen*95000 + remainingCars*10000)
}

func main() {
	productionRate := 200
	successRate := 85.0
	carsCount := 23

	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	workingCarsPerMinute := CalculateWorkingCarsPerMinute(productionRate, successRate)
	totalCost := CalculateCost(carsCount)

	fmt.Printf("Working cars per hour: %.2f\n", workingCarsPerHour)
	fmt.Printf("Working cars per minute: %d\n", workingCarsPerMinute)
	fmt.Printf("Total cost for %d cars: $%d\n", carsCount, totalCost)
}
