package main

import "fmt"

func main() {
	fmt.Println("Sum of even fibonacci numbers below 4,000,000")

	// create a slice to hold the fibonacci numbers, initialize with the first two (0, 1)
	fib := []int{0, 1}
	// sum := 0

	//calculate the fibonacci numbers below 4m
	for i := 2; ; i++ {
		// calculate the next number from the sum of the prior two
		next := fib[i-1] + fib[i-2]

		// break if the next number is larger than the limit
		if next >= 4_000_000 {
			break
		}

		// append the number
		fib = append(fib, next)
	}

	// concatenate the list and print it
	fmt.Println(fib)

}
