package dto

type CreateReviewRequest struct {
	Book int64  `json:"book"`
	Text string `json:"text"`
}

type CreateReviewResponse struct {
	ID string `json:"id"`
}

type GetReviewByIDResponse struct {
	ID   string `json:"id"`
	Book int64  `json:"book"`
	Text string `json:"text"`
}
