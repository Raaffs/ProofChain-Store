package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct{
	Client 		*mongo.Client
	Collection 	*mongo.Collection
}


func connectToMongoDB() (*mongo.Client, error) {
	if err:=godotenv.Load(".env");err!=nil{
		log.Fatal(err)
	}
    uri := os.Getenv("MONGO_URI")
	serverApi:=options.ServerAPI(options.ServerAPIVersion1)
    clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)
    
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return nil, err
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        return nil, err
    }
    log.Println("Connected to MongoDB Atlas!")
    return client, nil
}

func main(){
	client,err:=connectToMongoDB();if err!=nil{
		log.Fatal("Error connecting to mongodb: ",err)
	}

	app:=&App{
		Client: client,
	}
	app.Collection=app.Client.Database("ProofChain").Collection("Documents")
	log.Println("Collection selected:", app.Collection.Name())
	http.HandleFunc("POST /retrieve",app.RetrieveHandler)
	http.HandleFunc("POST /add",app.InsertHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Println("Error starting server:", err)
	}
}

func (app *App) RetrieveHandler(w http.ResponseWriter, r *http.Request) {
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
    document, err := app.Retrieve(sha.Sha);if err!=nil{
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

func(app *App)InsertHandler(w http.ResponseWriter,r *http.Request){
	var document Document

	err:=ReadJson(w,r,&document);if err!=nil{
		log.Println(err)
		WriteJson(w,http.StatusBadRequest,map[string]string{"error":"invalid json"})
		return
	}

	if err:=app.Insert(document);err!=nil{
		WriteJson(w,http.StatusInternalServerError,map[string]string{"Error":"Internal server error"})
		return
	}
	WriteJson(w,http.StatusOK,map[string]string{"Message":"Document inserted successfully"})
}
