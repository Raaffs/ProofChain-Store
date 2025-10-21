package mongorepo

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/Suy56/ProofChainStore/models"
)

type InstituteMongoRepository struct {
    collection *mongo.Collection
}

func NewInstituteMongoRepository(col *mongo.Collection) models.InstituteRepository {
    return &InstituteMongoRepository{collection: col}
}

func (r *InstituteMongoRepository) Insert(ctx context.Context, inst models.Institute) error {
    result, err := r.collection.InsertOne(ctx, inst)
    if err != nil {
        return err
    }
    log.Printf("Inserted institute with ID: %v\n", result.InsertedID)
    return nil
}

func (r *InstituteMongoRepository) RetrieveByName(ctx context.Context, name string) (*models.Institute, error) {
    var inst models.Institute
    filter := bson.D{{Key: "name", Value: name}}

    err := r.collection.FindOne(ctx, filter).Decode(&inst)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            log.Println("No institute found with name:", name)
            return nil, nil
        }
        return nil, err
    }
    return &inst, nil
}

func (r *InstituteMongoRepository) AddDocumentName(ctx context.Context, name, documentName string) error {
    filter := bson.D{{Key: "name", Value: name}}
    update := bson.D{{Key: "$addToSet", Value: bson.D{{Key: "documentNames", Value: documentName}}}} // avoids duplicates

    opts := options.Update().SetUpsert(true)
    _, err := r.collection.UpdateOne(ctx, filter, update, opts)
    if err != nil {
        log.Printf("Failed to add document name for %s: %v", name, err)
        return err
    }
    log.Printf("Added document '%s' to institute '%s'", documentName, name)
    return nil
}
