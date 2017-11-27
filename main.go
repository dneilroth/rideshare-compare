package main

import (
	"log"
	"runtime"

	"github.com/dneilroth/rideshare-compare/routers"
	"github.com/gorilla/handlers"
	"github.com/urfave/negroni"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := routers.InitRoutes()

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"HEAD", "OPTIONS", "POST", "GET"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.MaxAge(600),
	)

	n := negroni.Classic()

	n.UseHandler(cors(router))

	n.Run(":8080")
}
