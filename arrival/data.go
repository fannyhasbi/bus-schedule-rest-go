package arrival

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

func AddArrival(w http.ResponseWriter, r *http.Request) {
	var response ResponseAddArrival

	f := r.FormValue

	if !(len(f("id_perusahaan")) > 0 &&
		len(f("id_tujuan")) > 0 &&
		len(f("id_asal")) > 0 &&
		len(f("berangkat")) > 0 &&
		len(f("datang")) > 0) {

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
	datang := fmt.Sprintf("%s %s", now, f("datang"))

	query := fmt.Sprintf("INSERT INTO kedatangan VALUES (null, %d, %d, %d, '%s', '%s')", id_perusahaan, id_tujuan, id_asal, berangkat, datang)

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
