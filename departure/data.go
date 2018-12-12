package departure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

func AddDeparture(w http.ResponseWriter, r *http.Request) {
	var response ResponseAddDeparture

	f := r.FormValue // get form value

	if !(len(f("id_perusahaan")) > 0 &&
		len(f("id_tujuan")) > 0 &&
		len(f("id_asal")) > 0 &&
		len(f("berangkat")) > 0 &&
		len(f("sampai")) > 0) {

		response.Status = 400
		response.Message = "Bad Request"

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	db := data.Connect()
	defer db.Close()

	id_perusahaan, _ := strconv.Atoi(f("id_perusahaan"))
	id_tujuan, _ := strconv.Atoi(f("id_tujuan"))
	id_asal, _ := strconv.Atoi(f("id_asal"))

	t := time.Now()
	now := t.Format("2006-01-02")

	berangkat := fmt.Sprintf("%s %s", now, f("berangkat"))
	sampai := fmt.Sprintf("%s %s", now, f("sampai"))

	query := fmt.Sprintf("INSERT INTO keberangkatan VALUES (null, %d, %d, %d, '%s', '%s')", id_perusahaan, id_tujuan, id_asal, berangkat, sampai)

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
