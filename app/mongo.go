package app

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongoDB() (*mongo.Client, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	uri := os.Getenv("MONGO_URI")
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
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
