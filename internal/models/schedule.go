package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schedule struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	GameId          string             `json:"game_id" bson:"game_id"`
	IntervalMinutes int                `json:"interval_minutes" bson:"interval_minutes"`
	Times           []string           `json:"times" bson:"times"`
	Active          bool               `json:"active" bson:"active"`
	CreateAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdateAt        time.Time          `json:"updated_at" bson:"updated_at"`
}
