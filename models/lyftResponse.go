package models

import (
	"strconv"
)

// LyftAuthResponse ...
type LyftAuthResponse struct {
	Token string `json:"access_token"`
}

type lyftRideInfo struct {
	Type                string  `json:"ride_type"`
	EstimatedDuration   int     `json:"estimated_duration_seconds"`
	EstimatedDistance   float64 `json:"estimated_distance_miles"`
	EstimatedCostMax    int     `json:"estimated_cost_cents_max"`
	EstimatedCostMin    int     `json:"estimated_cost_cents_min"`
	PrimeTimePercantage string  `json:"primetime_percentage"`
	DisplayName         string  `json:"display_name"`
}

// LyftCostResponse ..
type LyftCostResponse struct {
	CostEstimates []lyftRideInfo `json:"cost_estimates"`
}

func (r lyftRideInfo) EstMaxCost() float64 {
	surge := r.PrimeTimePercantage
	surgeLen := len(surge)
	surge = surge[0 : surgeLen-1]
	surgeNum, _ := strconv.Atoi(surge)
	cost := float64(r.EstimatedCostMax) / 100.0
	if percentage := float64(surgeNum) / 100.0; percentage > 0 {
		cost = cost + (cost * percentage)
	}

	return cost
}

func (r lyftRideInfo) EstMinCost() float64 {
	return float64(r.EstimatedCostMin) / 100.0
}
