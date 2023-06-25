package model

// PricePlan defines the price plan metadata.
type PricePlan struct {
	EnergySupplier      string
	PlanName            string
	UnitRate            float64
	PeakTimeMultipliers []PeakTimeMultiplier
}
