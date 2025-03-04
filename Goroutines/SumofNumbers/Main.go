package main

import (
	"fmt"
	"time"
)

func sumofNumbers(start, end int, result *int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	*result = sum
}
func main() {
	var sum1, sum2 int
	go sumofNumbers(0, 50, &sum1)
	go sumofNumbers(51, 100, &sum2)
	time.Sleep(1 * time.Second)
	totalsum := sum1 + sum2
	fmt.Println("Total Sum of sum1 and sum 2 = ", totalsum)
}
