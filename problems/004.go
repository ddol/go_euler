package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	largest_palindrome := int64(0)

	// We assume the factors are 9xx to save compute time
	for i := 900; i < 1000; i++ {
		for j := 900; j < 1000; j++ {
			product := int64(i * j)

			// check if the product is a palindrome
			if utils.IsPalindromeNumber(product) {
				// check if the product is larger than the current largest palindrome
				if product > largest_palindrome {
					largest_palindrome = product
				}
			}
		}
	}

	fmt.Println("Largest palindromic product of two 3-digit numbers is: ", largest_palindrome)
}
