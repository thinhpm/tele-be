package handler

import (
	"be/internal/service"
)

type MongoDBHandler struct {
	service *service.MongoDBService
}

func NewMongoDBHandler(service *service.MongoDBService) *MongoDBHandler {
	return &MongoDBHandler{service: service}
}
