package controllers

import (
	"net/http"

	"github.com/newshipt/shipt-go-service-skeleton/version"
)

//PerformHealthCheckV1 returns a sample response and store query time
func PerformHealthCheckV1(w http.ResponseWriter, r *http.Request) {
	// TOOD: Tailor this portion of the healthcheck to the specific service using it
	// var healthCheck models.HealthCheck
	// var err error

	// vars := mux.Vars(r)
	// start := time.Now()

	// healthCheck.SearchSuggestions, err = models.GetSearchSuggestionsByFragment(vars["queryFragment"])

	// if err != nil {
	// 	info := make(map[string]string)
	// 	info["fragment"] = vars["queryFragment"]
	// 	errHandler.HandleErrorWithInfo(err, info)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// healthCheck.QueryTime = time.Since(start) / 1000

	// j, err := json.Marshal(&healthCheck)

	// if err != nil {
	// 	info := make(map[string]string)
	// 	info["fragment"] = vars["queryFragment"]
	// 	errHandler.HandleErrorWithInfo(err, info)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	j := []byte(`{"sample":"healthcheck"}`)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetVersionV1 Returns the current git commit hash
func GetVersionV1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(version.Version))
}
