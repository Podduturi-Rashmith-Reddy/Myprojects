package main

import "fmt"

// Function to calculate Fibonacci series
func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a // Send Fibonacci number to channel (blocks until receiver is ready)
		a, b = b, a+b
	}
	close(ch) // Close the channel when done
}

func main() {
	n := 10
	ch := make(chan int) // Unbuffered channel

	go fibonacci(n, ch) // Start the Fibonacci calculation in a goroutine

	// Receive and print Fibonacci numbers as they are sent
	for num := range ch {
		fmt.Println(num)
	}
}
