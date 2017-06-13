package server

import (
    "log"
    "net/http"
    "database/sql"

    "github.com/blankbook/shared/models"
    "github.com/blankbook/shared/web"
)

// SetupAPI adds the API routes to the provided router
func SetupAPI(r web.Router, db *sql.DB) {
    r.HandleRoute([]string{web.POST}, "/post", PostPost, db)
}

func PostPost(w http.ResponseWriter, queryParams map[string][]string, body string, db *sql.DB) {
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
    err = db.Ping()
    if err != nil {
        log.Fatalf(err.Error())
    }
    rows, err := db.Query("SELECT Title FROM posts")
    if err != nil {
        log.Fatalf(err.Error())
    }
    var s string
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&s)
        if err != nil {
            log.Fatal(err)
        }
        log.Println(s)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    // send post to database
    w.WriteHeader(http.StatusOK)
}
