package app

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Suy56/ProofChainStore/models"
	"github.com/Suy56/ProofChainStore/mongorepo"
	"github.com/Suy56/ProofChainStore/repository"
	"go.mongodb.org/mongo-driver/bson"
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

	db := client.Database("ProofChain") 

	return &App{
		Documents:  mongorepo.NewDocumentMongoRepository(db.Collection("Documents")),
		Institutes: mongorepo.NewInstituteMongoRepository(db.Collection("institute")),
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

// func (a *App) RetrieveDocumentHandler(w http.ResponseWriter, r *http.Request) {
// 	var payload struct {
// 		Shahash string `json:"shahash"`
// 	}
// 	if err := ReadJson(w, r, &payload); err != nil || payload.Shahash == "" {
// 		WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Missing shahash"})
// 		return
// 	}
// 	result, err := a.Documents.Retrieve(context.Background(), payload.Shahash)
// 	if err != nil {
// 		WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve document"})
// 		return
// 	}
// 	if result == nil {
// 		WriteJson(w, http.StatusNotFound, map[string]string{"error": "Document not found"})
// 		return
// 	}
// 	WriteJson(w, http.StatusOK, result)
// }

func (app *App) RetrieveDocumentHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RetrieveHandler called")
    sha := struct {
        Sha string `json:"shahash"`
    }{}
    err := json.NewDecoder(r.Body).Decode(&sha)
    if err != nil {
        log.Println("JSON decoding error:", err) // Log the error for debugging
        WriteJson(w, http.StatusBadRequest, map[string]string{"Error": "Invalid Json object"})
        return
    }

    document, err := app.Documents.Retrieve(context.Background(), sha.Sha);if err!=nil{
		log.Println("Error retrieving document:", err)
		WriteJson(w, http.StatusInternalServerError, map[string]string{"Error": "Failed to retrieve document"})
		return
	}

	jsonDoc:=struct{
		Shahash string `bson:"shahash" json:"shahash"`
		EncryptedDocument []byte `bson:"encryptedDocument" json:"encryptedDocument"`
		PublicAddress string `bson:"publicAddress" json:"publicAddress"`
	}{}

	marshalDoc,err:=bson.Marshal(document); if err!=nil{
		log.Println("Error : ",err)
        WriteJson(w, http.StatusInternalServerError, map[string]string{"Error": "Internal server error"})
        return
	}

	if err:=bson.Unmarshal(marshalDoc,&jsonDoc);err!=nil{
		log.Println("Error : ",err)
        WriteJson(w, http.StatusInternalServerError, map[string]string{"Error": "Internal server error"})
        return
	}
	log.Println("json doc:",jsonDoc.Shahash)
	
	WriteJson(w, http.StatusOK, jsonDoc)
}


// ----------------- Institute Handlers -----------------

func (app *App) InsertInstituteHandler(w http.ResponseWriter, r *http.Request) {
	var inst models.Institute
	if err := ReadJson(w, r, &inst); err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		return
	}
	if err := app.Institutes.Insert(context.Background(), inst); err != nil {
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
