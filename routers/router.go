package routers

import (
	"github.com/dneilroth/rideshare-compare/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//InitRoutes initialize all routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router = initRoutes(router)

	return router
}

func initRoutes(router *mux.Router) *mux.Router {
	compareRouter := mux.NewRouter()
	compareRouter.HandleFunc("/compare", controllers.Compare).Methods("GET")
	router.PathPrefix("/").Handler(negroni.New(negroni.Wrap(compareRouter)))

	return router
}
