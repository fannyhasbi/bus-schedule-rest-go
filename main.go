package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/fannyhasbi/bus-schedule-rest-go/arrival"
	"github.com/fannyhasbi/bus-schedule-rest-go/bus"
	"github.com/fannyhasbi/bus-schedule-rest-go/departure"
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
	router.HandleFunc("/api/add-bus", bus.AddBus).Methods("POST")

	router.HandleFunc("/api/departure", departure.ReturnDepartures).Methods("GET")
	router.HandleFunc("/api/add-departure", departure.AddDeparture).Methods("POST")

	router.HandleFunc("/api/arrival", arrival.ReturnArrivals).Methods("GET")
	router.HandleFunc("/api/add-arrival", arrival.AddArrival).Methods("POST")
	http.Handle("/", router)

	fmt.Printf("Connected to port %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
