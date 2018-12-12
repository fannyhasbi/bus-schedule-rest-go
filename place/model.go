package place

// Place struct represents single data of place
type Place struct {
	ID   int    `form:"id" json:"id"`
	Nama string `form:"nama" json:"nama"`
}

// ResponsePlace struct represents place data followed by HTTP status
type ResponsePlace struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Place `json:"data"`
}
