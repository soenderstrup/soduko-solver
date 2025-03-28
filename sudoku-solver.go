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
	fmt.Println(row, col)
	if col == len(sudoku[0]) {
		fmt.Println("done")
		return sudoku
	} else {
		nextRow, nextCol := getNextCell(sudoku, row, col)
		if sudoku[row][col] != 0 {
			return solveRecursively(sudoku, nextRow, nextCol)
		}
		for n := 1; n <= 9; n++{
			if valid(sudoku, row, col, n) {
				sudoku[row][col] = n
				return solveRecursively(sudoku, nextRow, nextCol)
			}
		}
		fmt.Println("could not solve")
		return sudoku
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
	solution := solve(sudoku)
	printSudoku(solution)
}