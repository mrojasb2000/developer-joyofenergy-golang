package model

// StoreReadings defines the ... metadata
type StoreReadings struct {
	SmartMeterId       string               `json:"smartMeterId"`
	ElectricityReading []ElectricityReading `json:"electricityReading"`
}
