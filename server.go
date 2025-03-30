package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello")
}

func solveSudoku(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data struct {
		Sudoku [][]int `json:"sudoku"`
	}
	err := decoder.Decode(&data)
	if err != nil || data.Sudoku == nil || len(data.Sudoku) != 9 {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
	solution := solve(data.Sudoku)
	json.NewEncoder(w).Encode(solution)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/solve", solveSudoku)
	PORT := "3000"

	fmt.Println("Server listening on: " + PORT)
	http.ListenAndServe("127.0.0.1:" + PORT, nil)
}