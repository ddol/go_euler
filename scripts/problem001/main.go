package main

import "fmt"

func main() {
	sum := 0
	max_number := 1000

	fmt.Println("The sum of all multiples of 3 or 5, below limit:", max_number)

	// iterate from 0 to the max number
	for i := 0; i < max_number; i++ {
		// check if the number modulus 3 or 5 is zero, if so there is no remainder
		// and the number is a multiple of 3 or 5
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println("The sum is: ", sum)
}
