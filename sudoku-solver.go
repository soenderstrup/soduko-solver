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

func solve(sudoku [][]int) {
	solveRecursively(sudoku, 0, 0)
}

func solveRecursively(sudoku [][]int, row int, col int) {
	if col == len(sudoku[0]) {
		printSudoku(sudoku)
		
	} else {
		nextRow, nextCol := getNextCell(sudoku, row, col)
		if sudoku[row][col] != 0 {
			solveRecursively(sudoku, nextRow, nextCol)
		} else {
			for n := 1; n <= 9; n++{
				if valid(sudoku, row, col, n) {
					sudoku[row][col] = n
					solveRecursively(sudoku, nextRow, nextCol)
					sudoku[row][col] = 0
				}
			}
		}
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
	solve(sudoku)
}