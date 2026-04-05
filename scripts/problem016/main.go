package main

import (
	"euler-go/pkg/utils"
	"fmt"
	"math/big"
)

func main() {
	exponent := 1000
	fmt.Println("Summing the digits of the number 2^", exponent)

	base := big.NewInt(2)

	result := new(big.Int).Exp(base, big.NewInt(int64(exponent)), nil)

	fmt.Printf("2^%d = %s\n", exponent, result.String())

	sum := utils.SumOfDigits(*result)

	fmt.Println("The sum of the digits is:", sum)
}
