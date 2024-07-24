package main

import (
	"fmt"
)

// Define the size of the Sudoku grid
const SIZE = 9

func main() {
	// Example Sudoku grid with 0 representing empty cells
	grid := [SIZE][SIZE]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println("Unsolved Sudoku")
	fmt.Printf("\n")

	printGrid(grid)
	fmt.Printf("\n\n")

	fmt.Println("Solved Sudoku")
	fmt.Printf("\n")

	if solveSudoku(&grid) {
		printGrid(grid)
	} else {
		fmt.Println("No solution exists")
	}

}

// Print the Sudoku grid
func printGrid(grid [SIZE][SIZE]int) {
	fmt.Println("-------------------------")
	for r := 0; r < SIZE; r++ {

		for d := 0; d < SIZE; d++ {
			if d == 0 {
				fmt.Printf("| ")
			}
			fmt.Print(grid[r][d], " ")
			if (d+1)%3 == 0 {
				fmt.Printf("| ")
			}
		}
		fmt.Println()
		if (r+1)%3 == 0 && r != SIZE-1 {
			fmt.Println("|-----------------------|")
		}
	}
	fmt.Println("-------------------------")
}

// Check if it is safe to place a number in the given position
func isSafe(grid [SIZE][SIZE]int, row int, col int, num int) bool {
	// Check if the number is already in the row
	for x := 0; x < SIZE; x++ {
		if grid[row][x] == num {
			return false
		}
	}

	// Check if the number is already in the column
	for x := 0; x < SIZE; x++ {
		if grid[x][col] == num {
			return false
		}
	}

	// Check if the number is in the 3x3 subgrid
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

// Solve the Sudoku using backtracking
func solveSudoku(grid *[SIZE][SIZE]int) bool {
	row := -1
	col := -1
	isEmpty := false

	// Find an empty cell in the grid
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if grid[i][j] == 0 {
				row = i
				col = j
				isEmpty = true
				break
			}
		}
		if isEmpty {
			break
		}
	}

	// No empty cell means the puzzle is solved
	if !isEmpty {
		return true
	}

	// Try placing numbers from 1 to 9 in the empty cell
	for num := 1; num <= 9; num++ {
		if isSafe(*grid, row, col, num) {
			grid[row][col] = num
			if solveSudoku(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}

	return false
}
