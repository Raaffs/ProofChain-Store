package main

import (
	"net/http"
    "log"
    "github.com/Suy56/ProofChainStore/app"
)

func main() {
    a := app.NewApp()

    // Document routes
    http.HandleFunc("POST /add", a.InsertDocumentHandler)
    http.HandleFunc("POST /retrieve", a.RetrieveDocumentHandler)

    // Institute routes
    http.HandleFunc("POST /institute/insert", a.InsertInstituteHandler)
    http.HandleFunc("GET /institute/retrieve", a.RetrieveInstituteHandler)
    http.HandleFunc("POST /institute/addDocument", a.AddDocumentToInstituteHandler)

    log.Println("Server started on :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
