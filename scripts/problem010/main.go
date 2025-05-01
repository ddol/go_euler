package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	limit := 2_000_000
	prime_sum := int64(0)

	fmt.Println("Calculating the sum of primes up to", limit)

	primes := utils.PrimesUpToLimit(int64(limit))

	for i := 0; i < len(primes); i++ {
		prime_sum += int64(primes[i])
	}

	fmt.Println("The sum is:", prime_sum)
}
