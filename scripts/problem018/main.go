package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func makeTriangle() [][]int {
	depth := 4

	triangle := make([][]int, depth)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
	}

	// initialise the triangle using slices for each row
	triangle[0] = []int{3}
	triangle[1] = []int{7, 4}
	triangle[2] = []int{2, 4, 6}
	triangle[3] = []int{8, 5, 9, 3}

	return triangle
}

func main() {
	triangle := makeTriangle()
	fmt.Println("This is going to be a problem, Number", utils.SumRange(4))
}
