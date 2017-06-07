package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/blankbook/shared/web"
        "github.com/blankbook/writecontent/server"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, write server here, I love %s!", r.URL.Path[1:])
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "The port to listen on")
	flag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	s := mux.NewRouter()
	web.SetupRoutes()
}
