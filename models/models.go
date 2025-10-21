package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Models struct {
	Documents  DocumentRepository
	Institutes InstituteRepository
}

type DocumentRepository interface {
    Insert(ctx context.Context, doc Document) error
    Retrieve(ctx context.Context, shahash string) (bson.M, error)
}

type InstituteRepository interface {
    Insert(ctx context.Context, inst Institute) error
    RetrieveByName(ctx context.Context, name string) (*Institute, error)
    AddDocumentName(ctx context.Context, name, documentName string) error
}

type Document struct {
    Shahash           string `bson:"shahash" json:"shahash"`
    EncryptedDocument []byte `bson:"encryptedDocument" json:"encryptedDocument"`
    PublicAddress     string `bson:"publicAddress" json:"publicAddress"`
}

type Institute struct {
    Name          string   `bson:"name" json:"name"`
    DocumentNames []string `bson:"documentNames" json:"documentNames"`
}
