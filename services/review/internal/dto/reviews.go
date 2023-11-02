package dto

type CreateReviewRequest struct {
	Text string `json:"text"`
}

type CreateReviewResponse struct {
	ID string `json:"id"`
}

type GetReviewByIDResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
