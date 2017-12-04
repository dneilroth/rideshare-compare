package uber

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dneilroth/rideshare-compare/models"

	"github.com/dneilroth/rideshare-compare/configuration"
	"github.com/dneilroth/rideshare-compare/repository"
)

// Estimate ... Uber estimate
func Estimate(pickupLat, pickupLong, dropoffLat, dropoffLong float64) models.CompareResponse {
	var cResp models.CompareResponse
	var respBody models.UberCostResponse

	config := configuration.GetNewAppConfig()
	url := config.UberURL
	endStr := "/v1.2/estimates/price?start_latitude=%f&start_longitude=%f&end_latitude=%f&end_longitude=%f"
	endpoint := fmt.Sprintf(endStr, pickupLat, pickupLong, dropoffLat, dropoffLong)
	req, err := http.NewRequest("GET", url+endpoint, nil)
	if err != nil {
		panic(err)
	}

	token := config.UberCredentials
	res, err := makeRequest(req, token)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(res.Body).Decode(&respBody)
	if err != nil {
		panic(err)
	}

	for _, est := range respBody.CostEstimates {
		if est.Type == "POOL" {
			cResp.Company = "Uber"
			cResp.Cost = est.EstimatedCostMax
		}
	}

	return cResp
}

func makeRequest(r *http.Request, token string) (*http.Response, error) {
	r.Header.Add("Authorization", "Token "+token)
	client := repository.NewHTTPClient()
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	return res, nil
}
