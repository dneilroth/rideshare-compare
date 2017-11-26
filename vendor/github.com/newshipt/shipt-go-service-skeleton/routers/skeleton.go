package routers

import (
	"github.com/gorilla/mux"
	"github.com/newshipt/shipt-go-service-skeleton/controllers"

	"github.com/urfave/negroni"
)

// SetSkeletonV1Routes exposes Sample routes in the router
func SetSkeletonV1Routes(router *mux.Router) *mux.Router {
	skeletonRouter := mux.NewRouter().PathPrefix("/v1").Subrouter().StrictSlash(false)

	skeletonRouter.HandleFunc("/skeleton-sample",
		controllers.GetSkeletonSample).Methods("GET")

	skeletonRouter.HandleFunc("/skeleton-sample/",
		controllers.GetSkeletonSample).Methods("GET")

	router.PathPrefix("/v1").Handler(negroni.New(
		negroni.Wrap(skeletonRouter),
	))

	return router
}
