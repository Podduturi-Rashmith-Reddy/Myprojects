package main

import (
	"fmt"
)

func sender(ch chan int) {
	// Send an integer to the receiver
	ch <- 42
}

func receiver(ch chan int, done chan bool) {
	// Receive the integer from the channel and print it
	num := <-ch
	fmt.Println("Received:", num)
	done <- true // Signal that the receiver is done
}

func main() {
	// Create channels for communication
	ch := make(chan int)
	done := make(chan bool) // Channel to signal completion

	// Start sender and receiver goroutines
	go sender(ch)
	go receiver(ch, done)

	// Wait for the receiver to finish before exiting
	<-done // Block until the receiver signals that it is done
}
