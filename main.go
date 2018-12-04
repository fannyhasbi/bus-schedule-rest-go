package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const PORT = ":8080"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	response := "This is web service for Bus Schedule React"
	fmt.Fprintln(w, response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/place", returnPlaces).Methods("GET")
	http.Handle("/", router)

	fmt.Printf("Connected to port %d", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
