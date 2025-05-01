package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	for i := 1; ; i++ {
		divisor_count := len(utils.Divisors(utils.SumRange(i)))
		if divisor_count > 500 {
			fmt.Println("The first triangular number with over 500 divisors is:", utils.SumRange(i))
			break
		}
	}
}
