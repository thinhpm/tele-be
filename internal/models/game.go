package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Account  []Account          `json:"account" bson:"account"`
	GameName string             `json:"game_name" bson:"game_name"`
	GameId   string             `json:"game_id" bson:"game_id"`
	CreateAt time.Time          `json:"created_at" bson:"created_at"`
	UpdateAt time.Time          `json:"updated_at" bson:"updated_at"`
}
