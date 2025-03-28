package main

import (
	"strconv"
	"strings"
	"testing"
)

func makeSudokus(input string) [][][]int {
	puzzles := strings.Split(strings.TrimSpace(input), "\n\n")
	sudokus := make([][][]int, len(puzzles))

	for p, puzzle := range puzzles {
		lines := strings.Split(strings.TrimSpace(puzzle), "\n")
		sudoku := make([][]int, 9)
		for i, line := range lines {
			line = strings.TrimSpace(line)
			sudoku[i] = make([]int, 9)
			for j, char := range line {
				number, err := strconv.Atoi(string(char))
				if err != nil {
					panic(err) // Handle error appropriately in real use cases
				}
				sudoku[i][j] = number
			}
		}
		sudokus[p] = sudoku
	}

	return sudokus
}

func TestSodukoSolver(t *testing.T) {

}