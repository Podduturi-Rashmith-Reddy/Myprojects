package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage generates a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border of stars above and below the message.
func AddBorder(message string, numStars int) string {
	border := strings.Repeat("*", numStars)
	return border + "\n" + message + "\n" + border
}

// CleanUpMessage removes stars and trims whitespace from a message.
func CleanUpMessage(message string) string {
	// Remove all '*'
	cleanedMessage := strings.ReplaceAll(message, "*", "")
	// Trim leading and trailing spaces/newlines
	return strings.TrimSpace(cleanedMessage)
}
func main() {

	fmt.Println(WelcomeMessage("Judy")) // Expected: "Welcome to the Tech Palace, JUDY"

	fmt.Println(AddBorder("Welcome!", 10))
	message := ``
	fmt.Println(CleanUpMessage(message)) // Expected: "BUY NOW, SAVE 10%"
}
