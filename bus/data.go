package bus

import (
	"encoding/json"
	"fmt"
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

func AddBus(w http.ResponseWriter, r *http.Request) {
	var response ResponseAddBus

	if len(r.FormValue("perusahaan")) == 0 {
		response.Status = 400
		response.Message = "Bad Request"

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	db := data.Connect()
	defer db.Close()

	perusahaan := r.FormValue("perusahaan")
	query := fmt.Sprintf("INSERT INTO perusahaan (nama) VALUES ('%s')", perusahaan)

	_, err := db.Exec(query)
	if err != nil {
		response.Status = 500
		response.Message = "Internal Server Error"

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println(err)
		return
	}

	response.Status = 200
	response.Message = "OK"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
