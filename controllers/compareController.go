package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dneilroth/rideshare-compare/google"
	"github.com/dneilroth/rideshare-compare/lyft"
	"github.com/dneilroth/rideshare-compare/models"
	"github.com/dneilroth/rideshare-compare/uber"
)

var authToken string

// Compare price of Uber vs Lyft for a given pickup + dropoff location
func Compare(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	pickup := strings.ToLower(vars["pickup"][0])
	dropoff := strings.ToLower(vars["dropoff"][0])
	pickupLoc := google.GetLatLong(pickup)
	dropoffLoc := google.GetLatLong(dropoff)
	pickupLat, pickupLong := pickupLoc.Lat, pickupLoc.Lng
	dropoffLat, dropoffLong := dropoffLoc.Lat, dropoffLoc.Lng

	estimates := []models.CompareResponse{}
	lyftEst := lyft.Estimate(pickupLat, pickupLong, dropoffLat, dropoffLong)
	estimates = append(estimates, lyftEst)
	uberEst := uber.Estimate(pickupLat, pickupLong, dropoffLat, dropoffLong)
	estimates = append(estimates, uberEst)

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
