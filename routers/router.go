package routers

import (
	"github.com/dneilroth/rideshare-compare/controllers"
	"github.com/gorilla/mux"
)

//InitRoutes initialize all routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router = initRoutes(router)

	return router
}

func initRoutes(router *mux.Router) *mux.Router {
	compareRouter := mux.NewRouter().PathPrefix("/").Subrouter()

	compareRouter.HandleFunc("/compare/",
		controllers.Compare).Methods("GET")

	return router
}
