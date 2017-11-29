package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dneilroth/rideshare-compare/lyft"
	"github.com/dneilroth/rideshare-compare/models"
)

var authToken string

// Compare price of Uber vs Lyft for a given pickup + dropoff location
func Compare(w http.ResponseWriter, r *http.Request) {
	var lyftEst models.CompareResponse

	c := make(chan models.CompareResponse)
	lyft.Estimate(c)
	estimates := []models.CompareResponse{}

	for i := 0; i < 1; i++ {
		estimates = append(estimates, <-c)
	}
	lyftEst = estimates[0]
	j, _ := json.Marshal(&lyftEst)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
