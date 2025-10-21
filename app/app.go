package app

import (
	"context"
	"log"
	"net/http"

	"github.com/Suy56/ProofChainStore/models"
	"github.com/Suy56/ProofChainStore/mongorepo"
	"github.com/Suy56/ProofChainStore/repository"
)


// ----------------- App -----------------

type App struct {
	Documents  repository.DocumentRepository
	Institutes repository.InstituteRepository
}

// ----------------- NewApp -----------------

func NewApp() *App {
	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	db := client.Database("mydb") // Hardcoded DB name, remove env dependency

	return &App{
		Documents:  mongorepo.NewDocumentMongoRepository(db.Collection("documents")),
		Institutes: mongorepo.NewInstituteMongoRepository(db.Collection("institutes")),
	}
}

// ----------------- Document Handlers -----------------

func (a *App) InsertDocumentHandler(w http.ResponseWriter, r *http.Request) {
	var doc models.Document
	if err := ReadJson(w, r, &doc); err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		return
	}
	if err := a.Documents.Insert(context.Background(), doc); err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to insert document"})
		return
	}
	WriteJson(w, http.StatusCreated, doc)
}

func (a *App) RetrieveDocumentHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Shahash string `json:"shahash"`
	}
	if err := ReadJson(w, r, &payload); err != nil || payload.Shahash == "" {
		WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Missing shahash"})
		return
	}
	result, err := a.Documents.Retrieve(context.Background(), payload.Shahash)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve document"})
		return
	}
	if result == nil {
		WriteJson(w, http.StatusNotFound, map[string]string{"error": "Document not found"})
		return
	}
	WriteJson(w, http.StatusOK, result)
}

// ----------------- Institute Handlers -----------------

func (a *App) InsertInstituteHandler(w http.ResponseWriter, r *http.Request) {
	var inst models.Institute
	if err := ReadJson(w, r, &inst); err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		return
	}
	if err := a.Institutes.Insert(context.Background(), inst); err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to insert institute"})
		return
	}
	WriteJson(w, http.StatusCreated, inst)
}

func (a *App) RetrieveInstituteHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string `json:"name"`
	}
	if err := ReadJson(w, r, &payload); err != nil || payload.Name == "" {
		WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Missing name"})
		return
	}
	result, err := a.Institutes.RetrieveByName(context.Background(), payload.Name)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve institute"})
		return
	}
	if result == nil {
		WriteJson(w, http.StatusNotFound, map[string]string{"error": "Institute not found"})
		return
	}
	WriteJson(w, http.StatusOK, result)
}

func (a *App) AddDocumentToInstituteHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name         string `json:"name"`
		DocumentName string `json:"documentName"`
	}
	if err := ReadJson(w, r, &payload); err != nil || payload.Name == "" || payload.DocumentName == "" {
		WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		return
	}
	if err := a.Institutes.AddDocumentName(context.Background(), payload.Name, payload.DocumentName); err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to add document to institute"})
		return
	}
	WriteJson(w, http.StatusOK, map[string]string{"message": "Document added successfully"})
}
