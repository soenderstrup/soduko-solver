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

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(strings.TrimSpace(string(input)), "\n") 

	sudokuBoard := make([][]int, 9)

	for i, line := range lines {
		line = strings.TrimSpace(line)
		sudokuBoard[i] = make([]int, 9)
		for j, char := range line {

			convertedNumber, err := strconv.Atoi(string(char))
			check(err)
			sudokuBoard[i][j] = convertedNumber
		}
	}

	for _, row := range sudokuBoard {
		fmt.Println(row)
	}
}