package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewsCollection struct {
	db *mongo.Collection
}

func (r *ReviewsCollection) Get() {
}
