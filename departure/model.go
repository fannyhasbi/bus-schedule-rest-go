package departure

type Departure struct {
	Id              int    `json:"id"`
	Id_Perusahaan   int    `json:"id_perusahaan"`
	Nama_Perusahaan string `json:"nama_perusahaan"`
	Id_Tujuan       int    `json:"id_tujuan"`
	Nama_Tujuan     string `json:"nama_tujuan"`
	Id_Asal         int    `json:"id_asal"`
	Nama_Asal       string `json:"nama_asal"`
	Berangkat       string `json:"berangkat"`
	Sampai          string `json:"sampai"`
}

type ResponseDeparture struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []Departure `json:"data"`
}

type ResponseAddDeparture struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
