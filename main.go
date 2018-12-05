package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/fannyhasbi/bus-schedule-rest-go/bus"
	"github.com/fannyhasbi/bus-schedule-rest-go/place"
)

const PORT = ":8080"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	response := "This is web service for Bus Schedule React"
	fmt.Fprintln(w, response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleIndex).Methods("GET", "POST")
	router.HandleFunc("/api/place", place.ReturnPlaces).Methods("GET")
	router.HandleFunc("/api/bus", bus.ReturnBuses).Methods("GET")
	http.Handle("/", router)

	fmt.Printf("Connected to port %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
