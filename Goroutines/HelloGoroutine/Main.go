package main

import (
	"fmt"
	"time"
)

func PrintHello() {
	fmt.Println("Hello GO routine")
}

func main() {
	go PrintHello()

	time.Sleep(1 * time.Second)
	fmt.Println("Hello Main")
}
