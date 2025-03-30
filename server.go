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
	var data map[string]any
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/solve", solveSudoku)

	http.ListenAndServe(":8000", nil)
}