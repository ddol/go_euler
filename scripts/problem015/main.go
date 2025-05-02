package main

import (
	"fmt"
)

func countPathsDP(n int) int64 {
	// Create a grid with (n+1) × (n+1) vertices
	grid := make([][]int64, n+1)
	for i := range grid {
		grid[i] = make([]int64, n+1)
	}

	// Initialize the rightmost column and bottom row to 1
	// because there's only one way to reach the end from these positions
	for i := 0; i <= n; i++ {
		grid[i][n] = 1 // Rightmost column
		grid[n][i] = 1 // Bottom row
	}

	// Fill the grid from bottom-right to top-left
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// Number of paths = paths going right + paths going down
			grid[i][j] = grid[i+1][j] + grid[i][j+1]
		}
	}

	// print the grid for debugging
	for _, row := range grid {
		fmt.Println(row)
	}

	// The answer is in the top-left cell
	return grid[0][0]
}

func main() {
	gridDim := 20
	fmt.Printf("Calculating all top left to bottom right paths on a %d x %d grid\n", gridDim, gridDim)

	paths := countPathsDP(gridDim)

	fmt.Println("Number of paths:", paths)
}
