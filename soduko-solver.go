package main

import (
	"fmt"
	"os"
) 

func main() {
	input, err := os.ReadFile("input.txt")
	if (err != nil) {
		panic(err)
	}

	fmt.Println(string(input))
}