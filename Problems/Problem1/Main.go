package main

import "fmt"

// CanFastAttack determines if a fast attack can be executed.
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy determines if spying is possible.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner determines if the prisoner can be signaled.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return prisonerIsAwake && !archerIsAwake
}

// CanFreePrisoner determines if the prisoner can be freed.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogPresent bool) bool {
	if petDogPresent {
		return !archerIsAwake
	}
	return prisonerIsAwake && !knightIsAwake && !archerIsAwake
}

func main() {
	// Test cases
	fmt.Println(CanFastAttack(true))  // Output: false
	fmt.Println(CanFastAttack(false)) // Output: true

	fmt.Println(CanSpy(false, true, false))  // Output: true
	fmt.Println(CanSpy(false, false, false)) // Output: false

	fmt.Println(CanSignalPrisoner(false, true)) // Output: true
	fmt.Println(CanSignalPrisoner(true, true))  // Output: false

	fmt.Println(CanFreePrisoner(false, false, true, false)) // Output: true
	fmt.Println(CanFreePrisoner(true, false, false, true))  // Output: true
	fmt.Println(CanFreePrisoner(true, true, false, false))  // Output: false
}
