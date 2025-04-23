package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	// big_prime := 600851475143
	big_prime := 13195

	fmt.Println("Calculating the largest prime factor of", big_prime)

	// create a slice to hold the prime factors
	// factors := []int{}

	// use the Primes function from the utils math package to generate a slice of primes up to the limit
	primes := utils.Primes(big_prime)

	// print the primes
	fmt.Println("Primes: ", primes)

}
