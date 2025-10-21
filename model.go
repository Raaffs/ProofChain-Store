package main

import (
	"context"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document struct {
    Shahash           string `bson:"shahash" json:"shahash"`
    EncryptedDocument []byte `bson:"encryptedDocument" json:"encryptedDocument"`
    PublicAddress     string `bson:"publicAddress" json:"publicAddress"`
}

func(app *App)Insert(document Document)error{
	result,err:=app.Collection.InsertOne(context.TODO(),document);if err!=nil{
		return err
	}
	log.Printf("Inserted document with ID: %v\n", result.InsertedID)
	return nil
}

func (app *App) Retrieve(shahash string) (bson.M, error) {
    var result bson.M // Use bson.M for structured output instead of `any`
    
    // Trim and sanitize input
    shahash = strings.TrimSpace(shahash)
    
    filter := bson.D{{Key: "shahash", Value: shahash}}
    log.Println("Query Filter:", filter)

    // Query the collection
    err := app.Collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            log.Println("No document found with shahash:", shahash)
            return nil, nil
        }
        log.Println("Internal server error:", err)
        return nil, err
    }    
    return result, nil
}
