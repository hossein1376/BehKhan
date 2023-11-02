package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hossein1376/BehKhan/review/internal/dto"
	"github.com/hossein1376/BehKhan/review/internal/repository/model"
	"github.com/hossein1376/BehKhan/review/pkg/transfer"
)

type ReviewsCollection struct {
	db *mongo.Collection
}

func (r *ReviewsCollection) Create(data *dto.CreateReviewRequest) (*dto.CreateReviewResponse, error) {
	result, err := r.db.InsertOne(context.Background(), data)
	if err != nil {
		return nil, transfer.InternalError{Err: err}
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, transfer.InternalError{Err: err}
	}

	return &dto.CreateReviewResponse{ID: id.Hex()}, nil
}

func (r *ReviewsCollection) Get(id string) (*dto.GetReviewByIDResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, transfer.BadRequestError{Err: err}
	}

	result := r.db.FindOne(context.Background(), bson.M{"_id": objectId}, nil)
	if result.Err() != nil {
		switch {
		case errors.Is(result.Err(), mongo.ErrNoDocuments):
			return nil, transfer.NotFoundError{Err: result.Err()}
		default:
			return nil, transfer.InternalError{Err: result.Err()}
		}
	}

	review := model.Review{}
	err = result.Decode(&review)
	if err != nil {
		return nil, transfer.InternalError{Err: err}
	}

	return &dto.GetReviewByIDResponse{ID: review.ID, Text: review.Text}, nil
}
