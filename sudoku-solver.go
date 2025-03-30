package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
) 


func solve(sudoku [][]int) [][]int {
	sudokuCopy := make([][]int, 9) // avoid side effects
	for i, row := range sudoku {
		sudokuCopy[i] = make([]int, 9)
		copy(sudokuCopy[i], row)
	}
	if solveRecursively(sudokuCopy, 0, 0) {
		return sudokuCopy
	} else {
		return nil
	}
}

func solveRecursively(sudoku [][]int, row int, col int) bool {
	if col == len(sudoku[0]) {
		return true
	} else {
		nextRow, nextCol := getNextCell(sudoku, row, col)
		if sudoku[row][col] != 0 {
			return solveRecursively(sudoku, nextRow, nextCol)
		} else {
			for n := 1; n <= 9; n++{
				if valid(sudoku, row, col, n) {
					sudoku[row][col] = n
					if solveRecursively(sudoku, nextRow, nextCol) {
						return true
					}
					sudoku[row][col] = 0
				}
			}
		}
		return false
	}
}

func getNextCell(sudoku [][]int, row int, col int) (int, int) {
	nextRow := row
	nextCol := col
	if row < len(sudoku) - 1 {
		nextRow++
	} else {
		nextRow = 0
		nextCol++
	}
	return nextRow, nextCol
}

func valid(sudoku [][]int, row int, col int, n int) bool {
	// I had to change this to make this function work to validate solutions. 
	// Not optimal...
	// for i := range len(sudoku) - 1 {
	// 	if n == sudoku[i][col] || n == sudoku[row][i] {
	// 		return false
	// 	}
	// }
	for i := range sudoku {
		if (i == col) {
			continue
		} else if n == sudoku[row][i] {
			return false
		}
	}
	for i := range sudoku {
		if (i == row) {
			continue
		} else if n == sudoku[i][col] {
			return false
		}
	}
	quadrantRow := int(math.Floor(float64(row) / 3) * 3)
	quadrantCol := int(math.Floor(float64(col) / 3) * 3)
	for i := quadrantRow; i < quadrantRow + 3; i++ {
		for j := quadrantCol; j < quadrantCol + 3; j++ {
			if i == row && j == col {
				continue
			} else if n == sudoku[i][j] {
				return false
			}
		}
	}
	return true
}

func validateSudoku(sudoku [][]int) bool {
	if sudoku == nil {
		return false
	}
	for i, row := range sudoku {
		for j, n := range row {
			if !valid(sudoku, i, j, n) {
				return false
			}
		}
	}
	return true
}

func printSudoku(sudoku [][]int) {
	for _, row := range sudoku {
		fmt.Println(row)
	}
}

func format2DSlice(slice [][]int) string {
	var sb strings.Builder
	for _, row := range slice {
		sb.WriteString("[ ")
		for _, n := range row {
			sb.WriteString(fmt.Sprintf("%d ", n))
		}
		sb.WriteString("]\n")
	}
	return sb.String()
}

func check(err error) {
	if (err != nil) {
		panic(err)
	}
}

func makeSudoku(input []byte) [][]int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n") 

	sudoku := make([][]int, 9)

	for i, line := range lines {
		line = strings.TrimSpace(line)
		sudoku[i] = make([]int, 9)
		for j, char := range line {
			number, err := strconv.Atoi(string(char))
			check(err)
			sudoku[i][j] = number
		}
	}
	return sudoku
}

func makeSudokus(input []byte) [][][]int {
	puzzles := strings.Split(strings.TrimSpace(strings.ReplaceAll(string(input), "\r", "")), "\n\n")
	sudokus := make([][][]int, len(puzzles))

	for p, puzzle := range puzzles {
		lines := strings.Split(strings.TrimSpace(puzzle), "\n")
		sudoku := make([][]int, 9)
		for i, line := range lines {
			line = strings.TrimSpace(line)
			sudoku[i] = make([]int, 9)
			for j, char := range line {
				number, err := strconv.Atoi(string(char))
				check(err)
				sudoku[i][j] = number
			}
		}
		sudokus[p] = sudoku
	}

	return sudokus
}
func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	sudoku := makeSudoku(input)
	printSudoku(sudoku)
	fmt.Println("")
	solution := solve(sudoku)
	if solution != nil {
		printSudoku(solution)
		fmt.Println(validateSudoku(solution))
	} else {
		fmt.Println("Solution not found.")
	}
}