package main

import "fmt"

func main() {
	count := 0
	for {
		fmt.Println("Running...")
		count++
		if count == 5 {
			break
		}
	}
}
