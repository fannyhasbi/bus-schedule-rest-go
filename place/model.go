package place

type Place struct {
	Id   int    `form:"id" json:"id"`
	Nama string `form:"nama" json:"nama"`
}

type ResponsePlace struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Place `json:"data"`
}
