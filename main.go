package main

import (
    "flag"
    "fmt"
    "net/http"

    "github.com/blankbook/writecontent/server"
)

func main() {
    server.SetupRoutes()

    var port int
    flag.IntVar(&port, "port", 80, "The port to listen on")
    flag.Parse()
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
