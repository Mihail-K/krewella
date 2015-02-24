/*
Command krewella is an HTTP POST -> IRC message gateway.
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"christine.website/go/krewella/irc"
	"github.com/codegangsta/negroni"
	"github.com/drone/routes"
)

var (
	router *routes.RouteMux
	n      *negroni.Negroni

	bots map[string]*irc.Bot
)

func createBots() error {
	bots = make(map[string]*irc.Bot)

	networklist := os.Getenv("KREWELLA_NETWORKS")
	networks := strings.Split(networklist, ",")

	for _, network := range networks {
		bot, err := irc.New(network)
		if err != nil {
			return err
		}

		bots[strings.ToLower(network)] = bot
	}

	return nil
}

func destroyBots() {
	for _, bot := range bots {
		bot.IrcObj.Quit()
	}
}

func main() {
	err := createBots()
	if err != nil {
		panic(err)
	}
	defer destroyBots()

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
