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
	depth := 15

	triangle := make([][]int, depth)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
	}

	// initialise the triangle using slices for each row
	triangle[0] = []int{75}
	triangle[1] = []int{95, 64}
	triangle[2] = []int{17, 47, 82}
	triangle[3] = []int{18, 35, 87, 10}
	triangle[4] = []int{20, 04, 82, 47, 65}
	triangle[5] = []int{19, 01, 23, 75, 03, 34}
	triangle[6] = []int{88, 02, 77, 73, 07, 63, 67}
	triangle[7] = []int{99, 65, 04, 28, 06, 16, 70, 92}
	triangle[8] = []int{41, 41, 26, 56, 83, 40, 80, 70, 33}
	triangle[9] = []int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29}
	triangle[10] = []int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14}
	triangle[11] = []int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57}
	triangle[12] = []int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48}
	triangle[13] = []int{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31}
	triangle[14] = []int{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23}

	return triangle
}

func main() {
	triangle := makeTriangle()
	maxSum := findMaxPathSum(triangle)
	fmt.Println("The maximum path sum is:", maxSum)
}
