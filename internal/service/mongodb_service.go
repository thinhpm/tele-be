package service

import (
	"be/internal/repository"
)

type MongoDBService struct {
	repo *repository.MongoDBRepository
}

func NewMongoDBService(repo *repository.MongoDBRepository) *MongoDBService {
	return &MongoDBService{repo: repo}
}
