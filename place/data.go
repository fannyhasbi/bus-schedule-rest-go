package place

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fannyhasbi/bus-schedule-rest-go/data"
)

// ReturnPlaces is handler function to return places
func ReturnPlaces(w http.ResponseWriter, r *http.Request) {
	var place Place
	var arrPlaces []Place
	var response ResponsePlace

	db := data.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tempat")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&place.ID, &place.Nama); err != nil {
			log.Fatal(err.Error())

		} else {
			arrPlaces = append(arrPlaces, place)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arrPlaces

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
