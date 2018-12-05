package bus

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fannyhasbi/bus-schedule-rest-go/data"
)

func ReturnBuses(w http.ResponseWriter, r *http.Request) {
	var bus Bus
	var arr_buses []Bus
	var response ResponseBus

	db := data.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM perusahaan")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&bus.Id, &bus.Nama); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_buses = append(arr_buses, bus)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arr_buses

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
