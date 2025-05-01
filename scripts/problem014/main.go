package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	fmt.Println("Identifying the longest Collatz sequence under 1 million")

	// Initialize variables to track the longest sequence and its starting number
	maxLength := 0
	maxStartingNumber := 0

	// Iterate through numbers from 1 to 1 million
	for i := 1; i < 1000000; i++ {
		// Calculate the Collatz sequence for the current number
		sequence := utils.CollatzSequence(i)

		// Get the length of the sequence
		length := len(sequence)

		// If the current sequence length is greater than the maximum found so far,
		// update the maximum length and starting number
		if length > maxLength {
			maxLength = length
			maxStartingNumber = i
		}
	}

	// Print the starting number that produces the longest Collatz sequence
	fmt.Println("Number", maxStartingNumber, "produces the longest Collatz sequence with length", maxLength)
}
