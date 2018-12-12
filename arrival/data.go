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

// ReturnArrivals is handler function to return arrival schedule
func ReturnArrivals(w http.ResponseWriter, r *http.Request) {
	var arrival Arrival
	var arrArrivals []Arrival
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
		if err := rows.Scan(&arrival.ID,
			&arrival.IDPerusahaan,
			&arrival.NamaPerusahaan,
			&arrival.IDTujuan,
			&arrival.NamaTujuan,
			&arrival.IDAsal,
			&arrival.NamaAsal,
			&arrival.Berangkat,
			&arrival.Sampai); err != nil {
			log.Fatal(err.Error())
		} else {
			arrArrivals = append(arrArrivals, arrival)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arrArrivals

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AddArrival is a handler function to add arrival schedule
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

	idPerusahaan, _ := strconv.Atoi(f("id_perusahaan"))
	idTujuan, _ := strconv.Atoi(f("id_tujuan"))
	idAsal, _ := strconv.Atoi(f("id_asal"))

	t := time.Now()
	now := t.Format("2006-01-02")

	berangkat := fmt.Sprintf("%s %s", now, f("berangkat"))
	datang := fmt.Sprintf("%s %s", now, f("datang"))

	query := fmt.Sprintf("INSERT INTO kedatangan VALUES (null, %d, %d, %d, '%s', '%s')", idPerusahaan, idTujuan, idAsal, berangkat, datang)

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
