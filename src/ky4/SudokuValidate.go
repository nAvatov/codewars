package ky4

import (
	"fmt"
)

func RunSudokuValidate() {
	var sudokuBoard = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	fmt.Println(ValidateSolution(sudokuBoard))
}

/*
4ky

Write a function validSolution/ValidateSolution/valid_solution() that accepts a 2D array representing a Sudoku board, and returns true if it is a valid solution, or false otherwise.
The cells of the sudoku board may also contain 0's, which will represent empty cells. Boards containing one or more zeroes are considered to be invalid solutions.

The board is always 9 cells by 9 cells, and every cell only contains integers from 0 to 9.

*/

func ValidateSolution(m [][]int) bool {
	var (
		linearPart []int
		k          int = 0
		g          int = 0
	)
	for k <= 6 {
		for g <= 6 {
			for i := k; i < k+3; i++ {
				for j := g; j < g+3; j++ {
					linearPart = append(linearPart, m[i][j])
				}
			}
			if !Validate(Sorted(linearPart)) {
				return false
			} else {
				linearPart = []int{}
			}
			g += 3
		}
		k += 3
		g = 0
	}
	return true
}

func Validate(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i] == 0 {
			return false
		} else {
			if row[i+1]-row[i] != 1 {
				return false
			}
		}
	}
	return true
}

func Sorted(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
