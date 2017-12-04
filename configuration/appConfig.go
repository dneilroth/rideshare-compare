package configuration

import (
	"os"
)

// AppConfig struct
type AppConfig struct {
	LyftURL          string
	UberURL          string
	LyftCredentials  string
	GoogleGeoCodeURL string
	UberCredentials  string
}

// GetNewAppConfig ...
func GetNewAppConfig() *AppConfig {
	lyftURL := "https://api.lyft.com"
	uberURL := "https://api.uber.com"
	lyftCredentials := os.Getenv("LYFT_CREDENTIALS")
	uberCredentials := os.Getenv("UBER_CREDENTIALS")
	googleGeoCodeURL := "https://maps.googleapis.com/maps/api/geocode/json"

	return &AppConfig{
		LyftURL:          lyftURL,
		UberURL:          uberURL,
		LyftCredentials:  lyftCredentials,
		GoogleGeoCodeURL: googleGeoCodeURL,
		UberCredentials:  uberCredentials,
	}
}
