package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/drone/routes"
)

var (
	router *routes.RouteMux
	n      *negroni.Negroni
)

func main() {
	router = routes.New()

	router.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "pancakes")
	})

	n = negroni.Classic()

	n.UseHandler(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	n.Run(":" + port)
}
