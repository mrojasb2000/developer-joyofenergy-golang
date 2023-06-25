package model

// Message defines the ... metadata
type Message struct {
	ID   string   `json:"id"`
	Data string   `json:"data"`
	Rows []string `json:"rows"`
}
