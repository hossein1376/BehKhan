package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hossein1376/BehKhan/review/internal/dto"
)

type Repository struct {
	Reviews ReviewRepository
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		Reviews: &ReviewsCollection{db: db},
	}
}

type ReviewRepository interface {
	Create(data *dto.CreateReviewRequest) (*dto.CreateReviewResponse, error)
	Get(bid int64, rid string) (*dto.GetReviewByIDResponse, error)
}
