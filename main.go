package main

import (
	"log"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/newshipt/shipt-wedge/configuration"
	"github.com/newshipt/shipt-wedge/routers"
	"github.com/urfave/negroni"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
	appConfig := configuration.GetAppConfig()

	router := routers.InitRoutes()

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"HEAD", "OPTIONS", "POST", "GET"}),
		handlers.AllowedHeaders([]string{
			"Content-Type",
			"X-Shipt-API-Token",
			"X-Shipt-API-Version",
			"X-Shipt-App-Version",
			"X-Shipt-Geo-Lat",
			"X-Shipt-Geo-Long",
		}),
		handlers.MaxAge(600),
	)

	n := negroni.Classic()

	n.UseHandler(cors(router))

	n.Run(appConfig.Server)
}
