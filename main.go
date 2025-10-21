package main

import (
	"net/http"
    "log"
    "github.com/Suy56/ProofChainStore/app"
)

func main() {
    app := app.NewApp()

    // Document routes
    http.HandleFunc("POST /add", app.InsertDocumentHandler)
    http.HandleFunc("POST /retrieve", app.RetrieveDocumentHandler)

    // Institute routes
    http.HandleFunc("POST /institute/insert", app.InsertInstituteHandler)
    http.HandleFunc("GET /institute/retrieve", app.RetrieveInstituteHandler)
    http.HandleFunc("POST /institute/addDocument", app.AddDocumentToInstituteHandler)

    log.Println("Server started on :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
