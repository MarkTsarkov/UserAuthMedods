package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    MongoDb := "mongodb://localhost:27017/"

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
    if err != nil {
        log.Fatal(err, "this field")
    }

    defer cancel()
    fmt.Println("Connected to MongoDB!")

    return client
}

//Client Database instance
var Client *mongo.Client = DBinstance()

//OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

    var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)

    return collection
}
