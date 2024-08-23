package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "accounts"

func (repo *MongoDBRepository) GetGameByGameId(ctx context.Context, gameId string) *mongo.SingleResult {
	collection := repo.db.Collection(collectionName)
	filter := bson.M{"game_id": gameId}
	return collection.FindOne(ctx, filter)
}

func (repo *MongoDBRepository) GetAllGames(ctx context.Context) (*mongo.Cursor, error) {
	collection := repo.db.Collection(collectionName)
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		log.Printf("Failed to find documents: %v", err)
		return nil, err
	}
	return cursor, nil
}
