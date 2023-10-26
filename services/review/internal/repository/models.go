package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Review struct {
	ID   string `bson:"_id"`
	Text string `bson:"text"`
}

type Repository struct {
	Reviews ReviewRepository
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		Reviews: &ReviewsCollection{db: db},
	}
}
