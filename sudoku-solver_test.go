package main

import (
	"os"
	"testing"
)

func TestSodukoSolver(t *testing.T) {
	input, err := os.ReadFile("input_test.txt")
	check(err)

	sudokus := makeSudokus(input)

	for _, sudoku := range sudokus {
		solution := solve(sudoku)
		if solution == nil {
			t.Errorf("Solution could not be found to:\n%s", format2DSlice(sudoku))
		} else if !validateSudoku(solution) {
			t.Errorf("Solution is not correct:\n%s", format2DSlice(solution))
		}
	}
}