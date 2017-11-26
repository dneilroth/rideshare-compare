package controllers

import "net/http"

//GetEmptyRequestV1 will be used to test the network speed
func GetEmptyRequestV1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}
