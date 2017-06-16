package server

import (
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
    var err error
    defer func() {
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
        }
    }()
    p, err := models.ParsePost(body)
    if err != nil {
        return
    }
    err = p.Validate()
    if err != nil {
        return
    }
    query :=`
    INSERT INTO Posts 
    (Title, Content, ContentType, GroupName, Time, Color)
    Values ($1, $2, $3, $4, $5, $6)`

    _, err = db.Exec(query, p.Title, p.Content, p.ContentType, p.GroupName, p.Time, p.Color)
    if err != nil {
        return
    }
    w.WriteHeader(http.StatusOK)
}
