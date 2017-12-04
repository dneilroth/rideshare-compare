package google

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/dneilroth/rideshare-compare/configuration"
	"github.com/dneilroth/rideshare-compare/models"
	"github.com/dneilroth/rideshare-compare/repository"
)

// GetLatLong ...
func GetLatLong(addr string) models.LatLng {
	var URL *url.URL

	config := configuration.GetNewAppConfig()
	geoCodeURL := config.GoogleGeoCodeURL
	googleAPIKey := os.Getenv("GOOGLE_API_KEY")
	URL, err := url.Parse(geoCodeURL)
	if err != nil {
		panic(err)
	}

	parameters := url.Values{}
	parameters.Add("address", addr)
	parameters.Add("key", googleAPIKey)

	URL.RawQuery = parameters.Encode()

	r, err := http.NewRequest("GET", URL.String(), nil)
	client := repository.NewHTTPClient()
	if err != nil {
		panic(err)
	}

	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var respBody models.GoogleResponse
	json.Unmarshal(body, &respBody)

	if len(respBody.Results) < 1 {
		return models.LatLng{}
	}

	return respBody.Results[0].Geometry.Location
}
