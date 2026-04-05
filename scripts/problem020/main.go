package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	input := 100
	fmt.Println("input number:", input)

	factorial := utils.Factorial(input)
	fmt.Println("factorial:", factorial)

	sum := utils.SumOfDigits(factorial)
	fmt.Println("sum of digits:", sum)
}
