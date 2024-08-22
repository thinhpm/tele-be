package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoDBService) GetGameByGameId(ctx context.Context, gameId string) *mongo.SingleResult {
	return s.repo.GetGameByGameId(ctx, gameId)
}

func (s *MongoDBService) GetAllGames(ctx context.Context) (*mongo.Cursor, error) {
	return s.repo.GetAllGames(ctx)
}
