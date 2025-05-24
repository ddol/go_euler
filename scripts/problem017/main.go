package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func main() {
	max_number := 1000
	fmt.Println("Sum of the letters in all English word representations of number, 1 to", max_number)

	sum := 0
	for i := 1; i <= max_number; i++ {
		word := utils.NumberToWordsNoSpaces(i)
		sum += word
	}

	fmt.Println("Sum:", sum)
}
