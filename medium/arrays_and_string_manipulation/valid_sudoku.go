package arrays_and_string_manipulation

import (
	"fmt"
	"strconv"
)

//Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
//Each row must contain the digits 1-9 without repetition.
//Each column must contain the digits 1-9 without repetition.
//Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//Note:
//
//A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//Only the filled cells need to be validated according to the mentioned rules.
//Input: board =
//[["5","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]
//Output: true

func isValidSudokuNaive(board [][]byte) bool {
	// Each row must contain the digits 1-9 without repetition.
	for i := 0; i < len(board); i++ {
		rowValues := map[string]bool{}
		for j := 0; j < len(board[0]); j++ {
			val := string(board[i][j])
			if rowValues[val] {
				// must be duplicate value in a row, that's not valid sudoku!
				fmt.Printf("duplicate row value found: %+v\n", val)
				return false
			} else if val != "." {
				rowValues[val] = true
			}
		}
	}

	// Each column must contain the digits 1-9 without repetition.
	for j := 0; j < len(board[0]); j++ {
		colValues := map[string]bool{}
		for i := 0; i < len(board); i++ {
			val := string(board[i][j])
			if colValues[val] {
				// must be duplicate value in a col, that's not valid sudoku!
				fmt.Printf("duplicate col value found: %+v\n", val)
				return false
			} else if val != "." {
				colValues[val] = true
			}
		}
	}

	// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
	subgridStartRow := 0
	subgridStartCol := 0
	for k := 0; k < 9; k++ {
		stopRow := subgridStartRow + 3
		subgridValues := map[string]bool{}
		for i := subgridStartRow; i < stopRow; i++ {
			stopCol := subgridStartCol + 3
			for j := subgridStartCol; j < stopCol; j++ {
				val := string(board[i][j])
				if subgridValues[val] {
					fmt.Printf("duplicate sub-grid value found: %+v\n", val)
					return false
				} else if val != "." {
					//fmt.Printf("Added value :%+v\n", val)
					subgridValues[val] = true
				}
			}
		}
		if subgridStartCol == 6 {
			// reset
			subgridStartCol = 0
		} else {
			subgridStartCol += 3
		}

		if k < 2 {
			subgridStartRow = 0
		} else if k >= 2 && k < 5 {
			subgridStartRow = 3
		} else {
			subgridStartRow = 6
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	seen := map[string]bool{}
	inRow := " in row "
	inCol := " in col "
	inBlock := " in block "

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			number := string(board[i][j])
			if number != "." {
				if seen[number+inRow+strconv.Itoa(i)] {
					return false
				} else {
					seen[number+inRow+strconv.Itoa(i)] = true
				}
				if seen[number+inCol+strconv.Itoa(j)] {
					return false
				} else {
					seen[number+inCol+strconv.Itoa(j)] = true
				}
				if seen[number+inBlock+strconv.Itoa(i/3)+"-"+strconv.Itoa(j/3)] {
					return false
				} else {
					seen[number+inBlock+strconv.Itoa(i/3)+"-"+strconv.Itoa(j/3)] = true
				}
			}
		}
	}
	return true
}
