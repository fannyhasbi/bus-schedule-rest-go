package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	response := "This is web service for Bus Schedule React"
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/", handleIndex)
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		fmt.Println("Error")
	}
}
