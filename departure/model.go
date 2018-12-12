package departure

// Departure struct represents single data of departure schedule
type Departure struct {
	ID             int    `json:"id"`
	IDPerusahaan   int    `json:"id_perusahaan"`
	NamaPerusahaan string `json:"nama_perusahaan"`
	IDTujuan       int    `json:"id_tujuan"`
	NamaTujuan     string `json:"nama_tujuan"`
	IDAsal         int    `json:"id_asal"`
	NamaAsal       string `json:"nama_asal"`
	Berangkat      string `json:"berangkat"`
	Sampai         string `json:"sampai"`
}

// ResponseDeparture struct represents departure data followed by HTTP status
type ResponseDeparture struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []Departure `json:"data"`
}

// ResponseAddDeparture struct represents null data followed by HTTP status
type ResponseAddDeparture struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
