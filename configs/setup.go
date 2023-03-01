package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// -------------------------------- mongoDB connection ----------------------------
func ConncectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to MongoDB")
	return client
}

//-------------------------------

var DB *mongo.Client = ConncectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("go_rest_api").Collection(collectionName)
}
