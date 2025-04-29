package utils

import "fmt"

// Generate a slice of primes up to a given limit using the Sieve of Eratosthenes
func Primes(limit int64) []int {
	// if the limit is less than 2, return an empty slice
	if limit < 2 {
		return []int{}
	}

	// create a slice to hold the primes
	primes := []int{}

	// generate a slice of booleans, initialized to true, up to the limit
	sieve := make([]bool, limit+1)

	// mark 0 and 1 as not prime
	sieve[0] = false
	sieve[1] = false

	// mark all other numbers as prime
	for i := 2; i <= int(limit); i++ {
		sieve[i] = true
	}

	// iterate over the sieve, starting at 2 and stopping at the square root of the limit
	for i := 2; i*i <= int(limit); i++ {
		// only continue if the value is still marked as prime
		if sieve[i] {
			// iterate over the multiples of the value, marking them as not prime
			for j := i * i; j <= int(limit); j += i {
				sieve[j] = false
			}
		}
	}

	// iterate over the sieve, if the value is true add the prime for that index to the primes slice
	for i := 2; i <= int(limit); i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func LargestPrimeFactor(n int64) int64 {
	// Special case for factor 2 (only even prime)
	for n%2 == 0 {
		n /= 2
	}

	// If n becomes 1, return 2 as the largest prime factor
	if n == 1 {
		return 2
	}

	// Check odd numbers only, starting from 3
	factor := int64(3)

	// Iterate through odd factors starting from 3
	// Stop when factor*factor exceeds n
	for factor*factor <= n {
		// If n is evenly divisible by factor (mod zero), divide n by factor and continue to check for larger factors
		if n%factor == 0 {
			n /= factor
		} else {
			// Skip even numbers
			factor += 2
		}
	}

	// If n > 1, n is the largest prime factor
	if n > 1 {
		return n
	}

	return factor
}

// Check if a given number is palindromic (i.e. 171)
func IsPalindromeNumber(n int64) bool {
	// convert int to string
	stringify_n := fmt.Sprintf("%d", n)

	// reverse string
	reverse_string := ""
	for i := len(stringify_n) - 1; i >= 0; i-- {
		reverse_string += string(stringify_n[i])
	}

	// are strings equal?
	return stringify_n == reverse_string
}
