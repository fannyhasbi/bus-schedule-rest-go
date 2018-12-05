package arrival

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fannyhasbi/bus-schedule-rest-go/data"
)

func ReturnArrivals(w http.ResponseWriter, r *http.Request) {
	var arrival Arrival
	var arr_arrivals []Arrival
	var response ResponseArrival

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
			k.datang
		FROM kedatangan k
		INNER JOIN perusahaan p
			ON k.id_perusahaan = p.id
		WHERE DAY(k.berangkat) = DAY(CURDATE())
		ORDER BY k.berangkat ASC;
	`)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&arrival.Id,
			&arrival.Id_Perusahaan,
			&arrival.Nama_Perusahaan,
			&arrival.Id_Tujuan,
			&arrival.Nama_Tujuan,
			&arrival.Id_Asal,
			&arrival.Nama_Asal,
			&arrival.Berangkat,
			&arrival.Sampai); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_arrivals = append(arr_arrivals, arrival)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arr_arrivals

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
