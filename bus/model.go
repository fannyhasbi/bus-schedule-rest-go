package bus

type Bus struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

type ResponseBus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Bus  `json:"data"`
}

type Perusahaan struct {
	Nama string `json:"perusahaan"`
}

type ResponseAddBus struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
