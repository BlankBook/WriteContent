package server

import (
    "net/http"

    "github.com/blankbook/shared/models"
    "github.com/blankbook/shared/web"
)

// SetupAPI adds the API routes to the provided router
func SetupAPI(r web.Router) {
    r.HandleRoute([]string{web.POST}, "/post", postPost)
}

func postPost(w http.ResponseWriter, queryParams map[string][]string, body string) {
    post, err := models.ParsePost(body)
    if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
    }
    err = post.Validate()
    if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
    }
    // send post to Cassandra database
    w.WriteHeader(http.StatusOK)
}
