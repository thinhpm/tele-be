package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoDBService) GetAllSchedules(ctx context.Context, active string) (*mongo.Cursor, error) {
	return s.repo.GetAllSchedules(ctx, active)
}
