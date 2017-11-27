package models

// LyftAuthResponse ...
type LyftAuthResponse struct {
	Token string `json:"access_token"`
}

type rideInfo struct {
	Type                string  `json:"ride_type"`
	EstimatedDuration   int     `json:"estimated_duration_seconds"`
	EstimatedDistance   float64 `json:"estimated_distance_miles"`
	EstimatedCostMax    int     `json:"estimated_cost_cents_max"`
	EstimatedCostMin    int     `json:"estimated_cost_cents_min"`
	PrimeTimePercantage string  `json:"primetime_percantage"`
	DisplayName         string  `json:"display_name"`
}

// LyftCostResponse ..
type LyftCostResponse struct {
	CostEstimates []rideInfo `json:"cost_estimates"`
}

func (r rideInfo) EstMaxCost() float64 {
	return float64(r.EstimatedCostMax) / 100.0
}

func (r rideInfo) EstMinCost() float64 {
	return float64(r.EstimatedCostMin) / 100.0
}
