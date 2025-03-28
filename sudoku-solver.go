package main

import (
	"fmt"
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
	solution := solveRecursively(sudoku, 0, 0)
	return solution
}

func solveRecursively(sudoku [][]int, row int, col int) [][]int {
	if col == len(sudoku[0]) {
		return sudoku
	} else {
		nextRow := row
		nextCol := col
		if row < len(sudoku) - 1 {
			nextRow++
		} else {
			nextRow = 0
			nextCol++
		}
		if sudoku[row][col] != 0 {
			return solveRecursively(sudoku, nextRow, nextCol)
		}
		for n := 1; n <= 9; n++{
			if legal(sudoku, row, col, n) {
				sudoku[row][col] = n
				return solveRecursively(sudoku, nextRow, nextCol)
			}
		}
		return sudoku
	}
}

func legal(sudoku[][]int, row int, col int, n int) bool {
	return true
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	sudoku := makeSudoku(input)
	solution := solve(sudoku)
	for _, row := range solution {
		fmt.Println(row)
	}
}