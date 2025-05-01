package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	prime_number := 10_001

	fmt.Println("Calculating primes up to prime number", prime_number)

	prime := utils.GenerateNthPrime(prime_number)

	fmt.Println("The", prime_number, "th prime number is", prime)
}
