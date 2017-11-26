package routers

import (
	"github.com/gorilla/mux"

	"github.com/newshipt/shipt-go-service-skeleton/controllers"
)

//SetMonitoringV1Routes Sets the routes for monitorization
func SetMonitoringV1Routes(router *mux.Router) *mux.Router {
	monitoringRouter := mux.NewRouter().PathPrefix("/v1/healthcheck").Subrouter()

	monitoringRouter.HandleFunc("/skeleton-sample",
		controllers.PerformHealthCheckV1).Methods("GET")

	monitoringRouter.HandleFunc("/skeleton-sample/",
		controllers.PerformHealthCheckV1).Methods("GET")

	monitoringRouter.HandleFunc("/network",
		controllers.GetEmptyRequestV1).Methods("GET")

	monitoringRouter.HandleFunc("/network/",
		controllers.GetEmptyRequestV1).Methods("GET")

	monitoringRouter.HandleFunc("/version",
		controllers.GetVersionV1).Methods("GET")

	monitoringRouter.HandleFunc("/version/",
		controllers.GetVersionV1).Methods("GET")

	router.PathPrefix("/v1/healthcheck").Handler(monitoringRouter)

	return router
}
