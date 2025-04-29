package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	smallest := int64(1)

	limit := 20

	fmt.Println("Finding the smallest number evenly divisible by all integers up to", limit)

	// moving in increments of 210 (2*3*5*7) efficiently find the smallest number that is evenly divisible by all numbers from 1 to limit
	for i := int64(210); ; i += 210 {
		if utils.DivisibleUpTo(i, limit) {
			smallest = i
			break
		}
	}

	fmt.Println("The smallest number evenly divisible by all integers up to", limit, "is", smallest)

}
