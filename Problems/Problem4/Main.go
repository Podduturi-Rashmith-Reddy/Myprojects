//Write a function WordCount that takes a string and returns a map
//where the keys are words and the values are the number of times each word appears.

package main

import (
	"fmt"
	"strings"
)

// WordCount function to count occurrences of each word
func WordCount(text string) map[string]int {
	wordCount := make(map[string]int)

	// Convert text to lowercase and split it into words
	words := strings.Fields(strings.ToLower(text))

	// Count occurrences of each word
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Go is great and Go is fun"
	result := WordCount(text)
	fmt.Println(result)
}
