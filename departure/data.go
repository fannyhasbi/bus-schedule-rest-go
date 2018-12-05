package departure

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fannyhasbi/bus-schedule-rest-go/data"
)

func ReturnDepartures(w http.ResponseWriter, r *http.Request) {
	var departure Departure
	var arr_departures []Departure
	var response ResponseDeparture

	db := data.Connect()
	defer db.Close()

	rows, err := db.Query(`
		SELECT k.id,
			k.id_perusahaan,
			p.nama AS nama_perusahaan,
			k.id_tujuan,
			(SELECT nama FROM tempat WHERE id = k.id_tujuan) AS nama_tujuan,
			k.id_asal,
			(SELECT nama FROM tempat WHERE id = k.id_asal) AS nama_asal,
			k.berangkat,
			k.sampai
		FROM keberangkatan k
		INNER JOIN perusahaan p
			ON k.id_perusahaan = p.id
		WHERE DAY(k.berangkat) = DAY(CURDATE())
		ORDER BY k.berangkat ASC;
	`)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&departure.Id,
			&departure.Id_Perusahaan,
			&departure.Nama_Perusahaan,
			&departure.Id_Tujuan,
			&departure.Nama_Tujuan,
			&departure.Id_Asal,
			&departure.Nama_Asal,
			&departure.Berangkat,
			&departure.Sampai); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_departures = append(arr_departures, departure)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arr_departures

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
