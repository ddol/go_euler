package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	upto := 10_000
	divSums := make([]int64, upto)
	amicables := []int64{}

	for i := 0; i < upto; i++ {
		divSums[i] = utils.SumList(utils.Divisors(int64(i + 1)))
	}

	for j := 0; j < upto; j++ {
		am1 := divSums[j]
		if am1 == 0 || am1 > int64(upto) || am1 == int64(j+1) {
			// move on if no proper divisors, out of range or self
			continue
		}
		am2 := divSums[am1-1]

		if int64(j+1) == am2 {
			amicables = append(amicables, am2)
		}
	}
	fmt.Println("amicables:", amicables)

	fmt.Println("sum:", utils.SumList(amicables))
}
