package models

// LatLng model ...
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// GeocodingResult is a single geocoded address
type GeocodingResult struct {
	FormattedAddress string          `json:"formatted_address"`
	Geometry         addressGeometry `json:"geometry"`
	Types            []string        `json:"types"`
	PlaceID          string          `json:"place_id"`
}

type addressGeometry struct {
	Location     LatLng `json:"location"`
	LocationType string `json:"location_type"`
}

// GoogleResponse model response from Google Geocoder API
type GoogleResponse struct {
	Status  string            `json:"status"`
	Results []GeocodingResult `json:"results"`
}
