package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	// big_prime := int64(13195)
	big_prime := int64(600851475143)

	fmt.Println("Calculating the largest prime factor of", big_prime)

	// print the last prime
	fmt.Println("Largest prime factor:", utils.LargestPrimeFactor(big_prime))
}
