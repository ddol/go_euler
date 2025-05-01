package main

import (
	"fmt"
)

func main() {
	result_a := 0
	result_b := 0
	result_c := 0

	fmt.Println("Calculating Pythagorean triplet (a^2 + b^2 = c^2) where a+b+c = 1000")

	for c := 0; c <= 1000; c++ {
		for b := 0; b < c; b++ {
			for a := 0; a < b; a++ {
				// Ensure the numbers sum to 1000
				if (a + b + c) == 1000 {
					// ensure the numbers are a pythagorean triplet
					if (a*a)+(b*b) == (c * c) {
						result_a = a
						result_b = b
						result_c = c
					}
				}
			}
		}
	}

	fmt.Println("Variables are a =", result_a, ", b =", result_b, ", c =", result_c)
	fmt.Println("Answer: ", result_a*result_b*result_c)
}
