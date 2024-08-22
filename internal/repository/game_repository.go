package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongoDBRepository) GetGameByGameId(ctx context.Context, gameId string) *mongo.SingleResult {
	filter := bson.M{"game_id": gameId}
	return repo.collection.FindOne(ctx, filter)
}

func (repo *MongoDBRepository) GetAllGames(ctx context.Context) (*mongo.Cursor, error) {
	filter := bson.M{}
	cursor, err := repo.collection.Find(ctx, filter)

	if err != nil {
		log.Printf("Failed to find documents: %v", err)
		return nil, err
	}
	return cursor, nil
}
