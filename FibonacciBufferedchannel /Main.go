package main

import "fmt"

// Function to calculate Fibonacci series
func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	n := 10
	ch := make(chan int, n) // Buffered channel to hold the Fibonacci numbers

	go fibonacci(n, ch) // Run Fibonacci function as a goroutine

	// Print the Fibonacci series
	for num := range ch {
		fmt.Println(num)
	}
}
