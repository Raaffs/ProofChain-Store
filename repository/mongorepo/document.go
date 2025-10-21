package mongorepo

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Suy56/ProofChainStore/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentMongoRepository struct {
    collection *mongo.Collection
}

func NewDocumentMongoRepository(col *mongo.Collection) models.DocumentRepository {
    return &DocumentMongoRepository{collection: col}
}

func (r *DocumentMongoRepository) Insert(ctx context.Context, doc models.Document) error {
    result, err := r.collection.InsertOne(ctx, doc)
    if err != nil {
        return err
    }
    log.Printf("Inserted document with ID: %v\n", result.InsertedID)
    return nil
}

func (r *DocumentMongoRepository) Retrieve(ctx context.Context, shahash string) (bson.M, error) {
    var result bson.M
    shahash = strings.TrimSpace(shahash)
    filter := bson.D{{Key: "shahash", Value: shahash}}

    err := r.collection.FindOne(ctx, filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            log.Println("No document found with shahash:", shahash)
            return nil, nil
        }
        return nil, err
    }
    fmt.Println(result)
    return result, nil
}
