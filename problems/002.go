package main

import "fmt"

func main() {
	sum := 0
	max_fibonacci := 4_000_000

	fmt.Println("Calculating the sum of even fibonacci numbers below", max_fibonacci)

	// create a slice to hold the fibonacci numbers, initialize with the first two (0, 1)
	fib := []int{0, 1}

	//calculate the fibonacci numbers below 4m
	for i := 2; ; i++ {
		// calculate the next number from the sum of the prior two
		next := fib[i-1] + fib[i-2]

		// break if the next number is larger than the limit
		if next >= max_fibonacci {
			break
		}

		// append the number
		fib = append(fib, next)
	}

	// iterate over the list, if the number is even, add it to the sum
	for _, v := range fib {
		if v%2 == 0 {
			sum += v
		}
	}

	// Print the sum
	fmt.Println("Sum: ", sum)

}
