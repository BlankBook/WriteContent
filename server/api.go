package server

import (
    "fmt"
    "net/http"

    "github.com/blankbook/shared/web"
)


func SetupAPI(r *web.HttpRouter) {
    r.HandleRoute(web.GET, "/test", getTest)
}

func getTest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Testing!!")
}
