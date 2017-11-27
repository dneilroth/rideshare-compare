package configuration

import "os"

// AppConfig struct
type AppConfig struct {
	LyftURL         string
	UberURL         string
	LyftCredentials string
}

// GetNewAppConfig ...
func GetNewAppConfig() *AppConfig {
	lyftURL := "https://api.lyft.com"
	uberURL := "https://api.uber.com"
	lyftCredentials := os.Getenv("LYFT_CREDENTIALS")

	return &AppConfig{
		LyftURL:         lyftURL,
		UberURL:         uberURL,
		LyftCredentials: lyftCredentials,
	}
}
