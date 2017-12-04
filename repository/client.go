package repository

import (
	"net/http"
	"time"
)

var client *http.Client

// NewHTTPClient ...
func NewHTTPClient() *http.Client {
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 5,
			Transport: &http.Transport{
				MaxIdleConns:        20,
				MaxIdleConnsPerHost: 20,
				DisableKeepAlives:   false,
			},
		}
	}

	return client
}
