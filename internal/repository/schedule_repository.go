package repository

import (
	"context"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongoDBRepository) GetAllSchedules(ctx context.Context, active string) (*mongo.Cursor, error) {
	collection := repo.db.Collection(collectionSchedules)

	filter := bson.M{}

	if active != "" {
		value, err := strconv.ParseBool(active)

		if err == nil {
			filter = bson.M{"active": value}
		}
	}

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		log.Println("Failed to find schedules: %v", err)
		return nil, err
	}

	return cursor, nil
}
