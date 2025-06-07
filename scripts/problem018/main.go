package main

import (
	"euler-go/pkg/utils"
	"fmt"
)

func findMaxPathSum(triangle [][]int) int {
	// Start from the second-to-last row and move upwards
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// Update the current cell with the max sum of the two adjacent cells below
			triangle[i][j] += utils.Max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	// The top element now contains the maximum path sum
	return triangle[0][0]
}

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
	maxSum := findMaxPathSum(triangle)
	fmt.Println("The maximum path sum is:", maxSum)
}
