package models

type Document struct {
    Shahash           string `bson:"shahash" json:"shahash"`
    EncryptedDocument []byte `bson:"encryptedDocument" json:"encryptedDocument"`
    PublicAddress     string `bson:"publicAddress" json:"publicAddress"`
}
