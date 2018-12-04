package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnPlaces(w http.ResponseWriter, r *http.Request) {
	var places Places
	var arr_place []Places
	var response ResponsePlaces

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tempat")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&places.Id, &places.Nama); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_place = append(arr_place, places)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_place

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
