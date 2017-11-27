package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dneilroth/rideshare-compare/configuration"
	"github.com/dneilroth/rideshare-compare/models"
	"github.com/dneilroth/rideshare-compare/repository"
)

var authToken string

// Compare price of Uber vs Lyft for a given pickup + dropoff location
func Compare(w http.ResponseWriter, r *http.Request) {
	config := configuration.GetNewAppConfig()
	client := repository.NewRideshareClient()
	url := config.LyftURL
	token := getAuthToken()
	endpoint := "/v1/cost?start_lat=37.7763&start_lng=-122.3918&end_lat=37.7972&end_lng=-122.4533"
	req, err := http.NewRequest("GET", url+endpoint, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "bearer "+token)
	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		panic(err)
	}
	var respBody models.LyftCostResponse
	err = json.NewDecoder(res.Body).Decode(&respBody)
	if err != nil {
		panic(err)
	}
	var cResp models.CompareResponse

	for _, est := range respBody.CostEstimates {
		if est.Type == "lyft_line" {
			cResp.Company = "Lyft"
			cResp.Cost = est.EstMaxCost()
		}
	}

	j, _ := json.Marshal(&cResp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getAuthToken() string {
	if authToken == "" {
		config := configuration.GetNewAppConfig()
		client := repository.NewRideshareClient()
		url := config.LyftURL
		data := config.LyftCredentials
		credsEnc := base64.StdEncoding.EncodeToString([]byte(data))
		jsonBody := []byte(`{"grant_type": "client_credentials", "scope": "public"}`)

		req, err := http.NewRequest("POST", url+"/oauth/token", bytes.NewBuffer(jsonBody))
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
