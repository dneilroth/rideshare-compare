package routers

import (
	"github.com/gorilla/mux"
)

//InitRoutes initialize all routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router = initV1Routes(router)

	return router
}

func initV1Routes(router *mux.Router) *mux.Router {
	router = SetSkeletonV1Routes(router)
	router = SetMonitoringV1Routes(router)
	return router
}
