package main

import (
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// TODO: instead of creating a webs server create things like a ping command

// startwebserver creates a http web server based on the input given
// the first argument will be the port used
// the second argument will be the router used
func startwebserver(args ...string) {
	if len(args) < 1 {
		color.Red("Not enough arguments to start a web server\n")
	}
	if len(args) == 1 {
		err := http.ListenAndServe(args[0], defaultserver())
		if err != nil {
			color.Red("something went wring when starting the webserver on args:%v", args[0])
		}
	}
}

func defaultserver() *mux.Router {
	router := mux.NewRouter()
	htmlServer := http.FileServer(http.Dir("/html/"))
	jsServer := http.FileServer(http.Dir("/js/"))
	cssServer := http.FileServer(http.Dir("/css/"))
	imgServer := http.FileServer(http.Dir("/images/"))
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Trying to serve file ./html/index.html")
		http.ServeFile(writer, request, "/html/index.html")
	})
	router.Handle("/html", htmlServer)
	router.Handle("/js", jsServer)
	router.Handle("/css", cssServer)
	router.Handle("/img", imgServer)
	return router
}
