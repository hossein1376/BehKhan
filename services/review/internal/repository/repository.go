package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hossein1376/BehKhan/review/internal/dto"
)

type Repository struct {
	Reviews ReviewRepository
}

func NewRepository(client *mongo.Client, db *mongo.Collection) *Repository {
	return &Repository{
		Reviews: &ReviewsCollection{client: client, db: db},
	}
}

type ReviewRepository interface {
	Create(data *dto.CreateReviewRequest) (*dto.CreateReviewResponse, error)
	Get(BookID int64, ReviewID string) (*dto.GetReviewByIDResponse, error)
}
