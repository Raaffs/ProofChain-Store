package main

import (
	"net/http"
    "log"
    "github.com/Suy56/ProofChainStore/app"
)
func logMethodMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("HTTP Method:", r.Method)
        next(w, r)
    }
}
func main() {
    app := app.NewApp()

    // Document routes
    http.HandleFunc("POST /add", logMethodMiddleware(app.InsertDocumentHandler))
    http.HandleFunc("GET /retrieve", logMethodMiddleware(app.RetrieveDocumentHandler))

    // Institute routes
    http.HandleFunc("POST /institute/insert", (app.InsertInstituteHandler))
    http.HandleFunc("GET /institute/retrieve", app.RetrieveInstituteHandler)
    http.HandleFunc("POST /institute/addDocument", app.AddDocumentToInstituteHandler)

    log.Println("Server started on :8754")
    log.Fatal(http.ListenAndServe(":8754", nil))
}
