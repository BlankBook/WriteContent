package server

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "database/sql"

    _ "github.com/denisenkom/go-mssqldb"

    "github.com/gorilla/mux"

    "github.com/blankbook/shared/web"
)

const pathPrefix = "/content/write"
const databaseUsernameEnvVar = "BB_CONTENT_DB_USERNAME"
const databasePasswordEnvVar = "BB_CONTENT_DB_PASSWORD"
const databaseServerEnvVar = "BB_CONTENT_DB_SERVER"
const databaseName = "blankbookcontent"

// SetupRoutes configures the service API endpoints
func SetupRoutes() {
    databaseUsername := os.Getenv(databaseUsernameEnvVar)
    databasePassword := os.Getenv(databasePasswordEnvVar)
    databaseServer := os.Getenv(databaseServerEnvVar)

    db, err := sql.Open("mssql",
                        fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
                                    databaseUsername,
                                    databasePassword,
                                    databaseServer,
                                    databaseName))
    if err != nil {
        log.Fatal(err)
    }

    muxRouter := mux.NewRouter()
    muxRouter.NotFoundHandler = http.HandlerFunc(notFoundHandler)
    r := web.NewHTTPRouter(muxRouter, pathPrefix)
    SetupAPI(r, db)
    r.StartListening()
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "endpoint not found", http.StatusNotFound)
}
