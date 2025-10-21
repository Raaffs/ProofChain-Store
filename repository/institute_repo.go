package repository

import (
    "context"
    "github.com/Suy56/ProofChainStore/models"
)

type InstituteRepository interface {
    Insert(ctx context.Context, inst models.Institute) error
    RetrieveByName(ctx context.Context, name string) (*models.Institute, error)
    AddDocumentName(ctx context.Context, name, documentName string) error
}
