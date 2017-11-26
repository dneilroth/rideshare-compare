package controllers

import (
	"net/http"
)

// GetSkeletonSample returns plain text skeleton
func GetSkeletonSample(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`

  .-=-=-.
 /       \
|         |
| )     ( |
\/ .\ /. \/
(    ^    )
 |.     .|
 ||xxxxx||
 '-._._.-'

	`))
}
