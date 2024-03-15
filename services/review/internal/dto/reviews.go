package dto

import (
	"github.com/hossein1376/BehKhan/review/pkg/domain"
)

type CreateReviewRequest struct {
	Book   int64         `json:"book"`
	Rating domain.Rating `json:"rating"`
	Text   string        `json:"text"`
}

type CreateReviewResponse struct {
	ID      string  `json:"id"`
	Average float64 `json:"-"`
	Total   int64   `json:"-"`
}

type GetReviewByIDResponse struct {
	ID   string `json:"id"`
	Book int64  `json:"book"`
	Text string `json:"text"`
}
