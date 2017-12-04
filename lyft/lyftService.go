package lyft

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dneilroth/rideshare-compare/configuration"
	"github.com/dneilroth/rideshare-compare/models"
	"github.com/dneilroth/rideshare-compare/repository"
)

var authToken string
var retries int

// Estimate get estimated lyft line cost for given
// pickup and dropoff location
func Estimate(pickupLat, pickupLong, dropoffLat, dropoffLong float64) models.CompareResponse {
	var cResp models.CompareResponse
	var respBody models.LyftCostResponse

	config := configuration.GetNewAppConfig()
	url := config.LyftURL
	endStr := "/v1/cost?start_lat=%f&start_lng=%f&end_lat=%f&end_lng=%f"
	endpoint := fmt.Sprintf(endStr, pickupLat, pickupLong, dropoffLat, dropoffLong)
	req, err := http.NewRequest("GET", url+endpoint, nil)
	if err != nil {
		panic(err)
	}

	res, err := makeRequest(req)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(res.Body).Decode(&respBody)
	if err != nil {
		panic(err)
	}

	for _, est := range respBody.CostEstimates {
		if est.Type == "lyft_line" {
			cResp.Company = "Lyft"
			cResp.Cost = est.EstMaxCost()
		}
	}

	return cResp
}

func getAuthToken() string {
	if authToken == "" {
		config := configuration.GetNewAppConfig()
		url := config.LyftURL
		data := config.LyftCredentials
		client := repository.NewHTTPClient()
		credsEnc := base64.StdEncoding.EncodeToString([]byte(data))
		jBody := []byte(`{"grant_type": "client_credentials", "scope": "public"}`)

		req, err := http.NewRequest("POST", url+"/oauth/token", bytes.NewBuffer(jBody))
		if err != nil {
			panic(err)
		}

		req.Header.Add("Authorization", "Basic "+credsEnc)
		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if res != nil {
			defer res.Body.Close()
		}
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var respBody models.LyftAuthResponse
		json.Unmarshal(body, &respBody)
		authToken = respBody.Token
	}

	return authToken
}

func makeRequest(r *http.Request) (*http.Response, error) {
	token := getAuthToken()
	r.Header.Add("Authorization", "bearer "+token)
	client := repository.NewHTTPClient()
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	if res.StatusCode == 401 {
		authToken = ""
		retries++
		if retries > 3 {
			retries = 0
			return res, errors.New("cannot authenticate with Lyft")
		}
		makeRequest(r)
	}
	retries = 0

	return res, nil
}
