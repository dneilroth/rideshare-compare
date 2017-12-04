package models

// CompareResponse ...
type CompareResponse struct {
	Company string  `json:"company"`
	Cost    float64 `json:"cost"`
}

// CompareJSON ...
type CompareJSON struct {
	Results []CompareResponse `json:"results"`
}
