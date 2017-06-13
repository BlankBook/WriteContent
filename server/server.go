package server

import (
    "os"

    "net/http"

    "github.com/gorilla/mux"

    "github.com/blankbook/shared/web"
)

const pathPrefix = "/content/write"

// SetupRoutes configures the service API endpoints
func SetupRoutes() {
   

    muxRouter := mux.NewRouter()
    muxRouter.NotFoundHandler = http.HandlerFunc(notFoundHandler)
    r := web.NewHTTPRouter(muxRouter, pathPrefix)
    SetupAPI(r)
    r.StartListening()
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "endpoint not found", http.StatusNotFound)
}
