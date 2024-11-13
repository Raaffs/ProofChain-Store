package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document struct {
    Shahash           string `bson:"shahash" json:"shahash"`
    EncryptedDocument string `bson:"encryptedDocument" json:"encryptedDocument"`
    PublicAddress     string `bson:"publicAddress" json:"publicAddress"`
}

func(app *App)Insert(document Document)error{
	result,err:=app.Collection.InsertOne(context.TODO(),document);if err!=nil{
		return err
	}
	log.Printf("Inserted document with ID: %v\n", result.InsertedID)
	return nil
}

func(app *App)Retrieve(shahash string)(any,error){
    var result any
    fmt.Println("shahash string",shahash)
	filter := bson.D{{Key: "shahash", Value: shahash}}
	err := app.Collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            log.Println("No document found with shahash:", result)
            return nil, nil
        }
        log.Println("internal server error:", result)
        return nil, err
    }
    fmt.Println("document : ",result)
    return result, nil
}	