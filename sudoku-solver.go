package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
) 

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
	for i := range len(sudoku) - 1 {
		if n == sudoku[i][col] || n == sudoku[row][i] {
			return false
		}
	}
	quadrantRow := int(math.Floor(float64(row) / 3) * 3)
	quadrantCol := int(math.Floor(float64(col) / 3) * 3)
	for i := quadrantRow; i < quadrantRow + 3; i++ {
		for j := quadrantCol; j < quadrantCol + 3; j++ {
			if i == row && j == col {
				continue
			}
			if n == sudoku[i][j] {
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

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	sudoku := makeSudoku(input)
	printSudoku(sudoku)
	fmt.Println("")
	solution := solve(sudoku)
	if solution != nil {
		printSudoku(solution)
	} else {
		fmt.Println("Solution not found.")
	}
}