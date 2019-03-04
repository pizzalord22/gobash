package main

import (
    "context"
    "fmt"
    "github.com/fatih/color"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "strconv"
    "time"
)

var startStop = make(chan string)

func webServer() {
    srv := &http.Server{
        Addr:         fmt.Sprintf("0.0.0.0:%s", "8080"),
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 15,
        Handler:      defaultRouter(),
    }
    for {
        select {
        case x := <-startStop:
            switch x {
            case "stop":
                err := srv.Shutdown(context.Background())
                CheckError(err, " webserver should have stopped gracefully")
            default:
                _, err := strconv.Atoi(x)
                if err != nil && len(x) < 5 {
                    color.Red("Could not start a webserver with %v as port\n switched to default port(8080)", x)
                    srv.Addr = "0.0.0.0:8080"
                }
                go func() {
                    err := srv.ListenAndServe()
                    if err != nil {
                        color.Red("something went wring when starting the webserver on port:%v", x)
                    }
                }()
            }
        }
    }
}

// startwebserver creates a http web server based on the input given
// the first argument will be the port used
func startwebserver(args ...string) {
    if len(args) < 1 {
        args = append(args, "8080")
    }
    startStop <- args[0]
}

// stopwebserver stops the webserver
func stopWebserver(w http.ResponseWriter, r *http.Request) {
    startStop <- "stop"
}

// defaultROuter is used a a base for possible future router
func defaultRouter() *mux.Router {
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
    router.HandleFunc("/api/v1/webserver/stop", stopWebserver)
    return router
}
