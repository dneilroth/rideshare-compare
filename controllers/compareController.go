package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dneilroth/rideshare-compare/google"
	"github.com/dneilroth/rideshare-compare/lyft"
	"github.com/dneilroth/rideshare-compare/models"
)

var authToken string

// Compare price of Uber vs Lyft for a given pickup + dropoff location
func Compare(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	pickup := strings.ToLower(vars["pickup"][0])
	dropoff := strings.ToLower(vars["dropoff"][0])
	pickupChan := make(chan models.LatLng)
	defer close(pickupChan)
	dropoffChan := make(chan models.LatLng)
	defer close(dropoffChan)

	go func() { pickupChan <- google.GetLatLong(pickup) }()
	go func() { dropoffChan <- google.GetLatLong(dropoff) }()

	pickupLoc := <-pickupChan
	dropoffLoc := <-dropoffChan
	pickupLat, pickupLong := pickupLoc.Lat, pickupLoc.Lng
	dropoffLat, dropoffLong := dropoffLoc.Lat, dropoffLoc.Lng

	c := make(chan models.CompareResponse)
	defer close(c)
	go func() { c <- lyft.Estimate(pickupLat, pickupLong, dropoffLat, dropoffLong) }()
	go func() { c <- lyft.Estimate(pickupLat, pickupLong, dropoffLat, dropoffLong) }()
	estimates := []models.CompareResponse{}

	for i := 0; i < 2; i++ {
		estimate := <-c
		estimates = append(estimates, estimate)
	}
	var resp models.CompareJSON
	for _, est := range estimates {
		resp.Results = append(resp.Results, est)
	}

	j, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
