package repository

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/Suy56/ProofChainStore/models"
)

type DocumentRepository interface {
    Insert(ctx context.Context, doc models.Document) error
    Retrieve(ctx context.Context, shahash string) (bson.M, error)
}
