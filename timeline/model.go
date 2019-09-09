package  timeline

type Products struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
type Timeline struct {
	Timestamp     string     `json:"timestamp"`
	Revenue       float64    `json:"revenue"`
	TransactionID string     `json:"transaction_id"`
	StoreName     string     `json:"store_name"`
	Products      []Products `json:"products"`
}


type Events struct {
	Events []Event `json:"events"`
}

type CustomData struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
type Event struct {
	Event      string       `json:"event"`
	Timestamp  string       `json:"timestamp"`
	CustomData []CustomData `json:"custom_data"`
	Revenue    int          `json:"revenue,omitempty"`
}