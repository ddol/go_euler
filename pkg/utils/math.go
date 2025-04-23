package utils

// Generate a slice of primes up to a given limit using the Sieve of Eratosthenes
func Primes(limit int) []int {
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
	for i := 2; i <= limit; i++ {
		sieve[i] = true
	}

	// iterate over the sieve, if the value is true add the prime for that index to the primes slice
	for i := 2; i <= limit; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
