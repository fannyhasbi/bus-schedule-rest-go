package main

type Places struct {
	Id   int    `form:"id" json:"id"`
	Nama string `form:"nama" json:"nama"`
}

type ResponsePlaces struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Places `json:"data"`
}
