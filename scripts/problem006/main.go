package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	limit := 100
	fmt.Println("Difference between sum of squares and square of sum of numbers up to", limit)

	result := utils.SquareOfSum(limit) - utils.SumOfSquare(limit)

	fmt.Println("The result is:", result)
}
