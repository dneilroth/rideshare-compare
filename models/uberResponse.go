package models

type uberRideInfo struct {
	Type              string  `json:"display_name"`
	EstimatedDuration int     `json:"duration"`
	EstimatedDistance float64 `json:"distance"`
	EstimatedCostMax  float64 `json:"high_estimate"`
	EstimatedCostMin  float64 `json:"low_estimate"`
}

// UberCostResponse ..
type UberCostResponse struct {
	CostEstimates []uberRideInfo `json:"prices"`
}
