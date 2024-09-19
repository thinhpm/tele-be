package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongoDBRepository) GetGameByGameId(ctx context.Context, gameId string) *mongo.SingleResult {
	collection := repo.db.Collection(collectionGames)
	filter := bson.M{"game_id": gameId}
	return collection.FindOne(ctx, filter)
}

func (repo *MongoDBRepository) GetAllGames(ctx context.Context) (*mongo.Cursor, error) {
	collection := repo.db.Collection(collectionGames)
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		log.Printf("Failed to find documents: %v", err)
		return nil, err
	}
	return cursor, nil
}

func (repo *MongoDBRepository) UpdateGameByGameId(ctx context.Context, gameId string, updates bson.M) (*mongo.UpdateResult, error) {
    collection := repo.db.Collection(collectionGames)
	filter := bson.M{"game_id": gameId}
    update := bson.M{"$set": updates}

	result, err := collection.UpdateOne(ctx, filter, update)

    if err != nil {
		log.Printf("Failed to update document with game_id %s: %v", gameId, err)
		return nil, err
	}

	return result, nil
}