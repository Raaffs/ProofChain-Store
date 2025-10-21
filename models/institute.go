package models

type Institute struct {
    Name          string   `bson:"name" json:"name"`
    DocumentNames []string `bson:"documentNames" json:"documentNames"`
}
