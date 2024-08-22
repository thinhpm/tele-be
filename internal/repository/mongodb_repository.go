package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoDBRepository(uri, dbName, collectionName string) (*MongoDBRepository, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, err
	}
	// log.Fatalln(dbName, collectionName)

	collection := client.Database(dbName).Collection(collectionName)
	// log.Fatalln(collection)

	return &MongoDBRepository{
		client:     client,
		collection: collection,
	}, nil
}

func (repo *MongoDBRepository) Disconnect() error {
	err := repo.client.Disconnect(context.TODO())

	if err != nil {
		log.Printf("Failed to disconnect from MongoDB: %v", err)
	}

	return err
}
