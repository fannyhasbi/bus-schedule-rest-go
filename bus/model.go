package bus

// Bus struct represents single bus data
type Bus struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

// ResponseBus struct represents bus data followed by HTTP status
type ResponseBus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Bus  `json:"data"`
}

// ResponseAddBus struct represents null data followed by HTTP status
type ResponseAddBus struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
