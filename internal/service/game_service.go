package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoDBService) GetGameByGameId(ctx context.Context, gameId string) *mongo.SingleResult {
	return s.repo.GetGameByGameId(ctx, gameId)
}

func (s *MongoDBService) GetAllGames(ctx context.Context) (*mongo.Cursor, error) {
	return s.repo.GetAllGames(ctx)
}

func (s *MongoDBService) UpdateGameByGameId(ctx context.Context, gameId string, updates bson.M) (*mongo.UpdateResult, error) {
	return s.repo.UpdateGameByGameId(ctx, gameId, updates)
}
