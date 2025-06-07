package utils

import "fmt"

// Define grid structure
type Grid struct {
	Data       [][]int
	rows, cols int
}

// Generate a slice of primes up to a given limit using the Sieve of Eratosthenes
func PrimesUpToLimit(limit int64) []int {
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

func GenerateNthPrime(n int) int {
	// if n is less than 1, return 0
	if n < 1 {
		return 0
	}

	// create a slice to hold the primes
	primes := []int{}

	// start with the first prime number
	current := 2

	// iterate until we have found n primes
	for len(primes) < n {
		// check if current is prime
		isPrime := true
		for _, prime := range primes {
			if current%prime == 0 {
				isPrime = false
				break
			}
		}

		// if it is prime, add it to the slice
		if isPrime {
			primes = append(primes, current)
		}

		current++
	}

	return primes[n-1]
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

// Check if n is divisible by all numbers from 1 to limit
func DivisibleUpTo(n int64, limit int) bool {
	for i := 1; i <= limit; i++ {
		if n%int64(i) != 0 {
			return false
		}
	}
	return true
}

// Return the sum of all numbers in the range 1..n
func SumRange(n int) int64 {
	result := int64(0)

	for i := int64(1); i <= int64(n); i++ {
		result += i
	}

	return result
}

// square each number in the range 1..n and return their sum
func SumOfSquare(n int) int64 {
	result := int64(0)

	for i := int64(1); i <= int64(n); i++ {
		result += (i * i)
	}

	return result
}

// sum the range of numbers in the range 1..n and square the sum
func SquareOfSum(n int) int64 {
	sum := SumRange(n)

	return (sum * sum)
}

func Divisors(n int64) []int64 {
	divisors := []int64{}

	// find divisors of n
	for i := int64(1); i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

// Create a new grid with the given number of rows and columns
func NewGrid(rows, cols int) *Grid {
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols)
	}
	return &Grid{Data: data, rows: rows, cols: cols}
}

// find the maximum product of n adjacent numbers in the grid
func (g *Grid) MaxGridProduct(n int) int {
	maxProduct := 0

	// Check rows
	for i := 0; i < g.rows; i++ {
		for j := 0; j <= g.cols-n; j++ {
			product := 1
			for k := 0; k < n; k++ {
				product *= g.Data[i][j+k]
			}
			if product > maxProduct {
				maxProduct = product
			}
		}
	}

	// Check columns
	for i := 0; i <= g.rows-n; i++ {
		for j := 0; j < g.cols; j++ {
			product := 1
			for k := 0; k < n; k++ {
				product *= g.Data[i+k][j]
			}
			if product > maxProduct {
				maxProduct = product
			}
		}
	}

	// Check diagonals (top-left to bottom-right)
	for i := 0; i <= g.rows-n; i++ {
		for j := 0; j <= g.cols-n; j++ {
			product := 1
			for k := 0; k < n; k++ {
				product *= g.Data[i+k][j+k]
			}
			if product > maxProduct {
				maxProduct = product
			}
		}
	}

	// Check diagonals (top-right to bottom-left)
	for i := 0; i <= g.rows-n; i++ {
		for j := n - 1; j < g.cols; j++ {
			product := 1
			for k := 0; k < n; k++ {
				product *= g.Data[i+k][j-k]
			}
			if product > maxProduct {
				maxProduct = product
			}
		}
	}

	// Return the maximum product found
	return maxProduct
}

// Given an integer n, return the next number in the Collatz sequence
func NextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

// Given an integer n, return the Collatz sequence starting from n
func CollatzSequence(n int) []int {
	sequence := []int{n}

	for n != 1 {
		n = NextCollatz(n)
		sequence = append(sequence, n)
	}

	return sequence
}

// Given an integer n, return the number of letters in its British English word representation without spaces or hyphens
// e.g. 342 = "threehundredandfortytwo" = 23 letters
func NumberToWordsNoSpaces(n int) int {
	if n == 0 {
		return 0
	}

	ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	teens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	thousands := []string{"", "thousand"}

	words := ""

	if n >= 1000 {
		words += ones[n/1000] + thousands[1]
		n %= 1000
	}

	if n >= 100 {
		words += ones[n/100] + "hundred"
		if n%100 != 0 {
			words += "and"
		}
		n %= 100
	}

	if n >= 20 {
		words += tens[n/10]
		n %= 10
	}

	if n >= 10 {
		words += teens[n-10]
		n = 0
	}

	if n > 0 {
		words += ones[n]
	}

	return len(words)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
