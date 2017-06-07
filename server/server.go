package server

import (
    "github.com/gorilla/mux"

    "github.com/blankbook/shared/web"
)

const PATH_PREFIX = "/content/write"

func SetupRoutes() {
    muxRouter := mux.NewRouter()
    r := web.NewHttpRouter(muxRouter, PATH_PREFIX)
    SetupAPI(r)
    r.StartListening();
}
