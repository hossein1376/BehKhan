package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hossein1376/BehKhan/review/internal/dto"
	"github.com/hossein1376/BehKhan/review/internal/repository/model"
	"github.com/hossein1376/BehKhan/review/pkg/transfer"
)

type ReviewsCollection struct {
	client *mongo.Client
	db     *mongo.Collection
}

func (r *ReviewsCollection) Create(data *dto.CreateReviewRequest) (*dto.CreateReviewResponse, error) {
	result, err := r.db.InsertOne(context.Background(), data)
	if err != nil {
		return nil, transfer.InternalError{Err: fmt.Errorf("insert new review: %w", err)}
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, transfer.InternalError{Err: fmt.Errorf("read inserted review id: %w", err)}
	}

	average, total, err := r.aggregateAvgCount(data.Book)
	if err != nil {
		return nil, err
	}

	return &dto.CreateReviewResponse{ID: id.Hex(), Average: average, Total: total}, nil
}

func (r *ReviewsCollection) Get(BookID int64, ReviewID string) (*dto.GetReviewByIDResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(ReviewID)
	if err != nil {
		return nil, transfer.BadRequestError{Err: err}
	}

	result := r.db.FindOne(context.Background(), bson.M{"_id": objectId, "book": BookID}, nil)
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

func (r *ReviewsCollection) aggregateAvgCount(bookID int64) (float64, int64, error) {
	cursor, err := r.db.Aggregate(context.Background(), mongo.Pipeline{
		bson.D{{"$match", bson.D{{"book", bookID}}}},
		bson.D{{"$group", bson.D{
			{"_id", "$book"},
			{"average", bson.D{{"$avg", "$rating"}}},
			{"total", bson.D{{"$sum", 1}}},
		}}},
	})
	if err != nil {
		return 0, 0, transfer.InternalError{Err: fmt.Errorf("aggregating reviews: %w", err)}
	}

	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		return 0, 0, transfer.InternalError{Err: fmt.Errorf("reading aggregated reviews response: %w", err)}
	}
	if len(results) != 1 {
		return 0, 0, transfer.InternalError{Err: fmt.Errorf("number of aggregated responses: %d", len(results))}
	}

	average, ok := results[0]["average"].(float64)
	if !ok {
		return 0, 0, transfer.InternalError{Err: fmt.Errorf("type of 'average' aggregate response: %T", results[0]["average"])}
	}
	total, ok := results[0]["total"].(int32)
	if !ok {
		return 0, 0, transfer.InternalError{Err: fmt.Errorf("type of 'total' aggregate response: %T", results[0]["total"])}
	}

	return average, int64(total), nil
}
