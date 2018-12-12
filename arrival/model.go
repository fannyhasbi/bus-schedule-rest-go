package arrival

// Arrival struct represents single data of arrival schedule
type Arrival struct {
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

// ResponseArrival struct represents arrival data followed by HTTP status
type ResponseArrival struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Arrival `json:"data"`
}

// ResponseAddArrival struct represents null data followed by HTTP status
type ResponseAddArrival struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
