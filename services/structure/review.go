package structure

type UpdateBookReviews struct {
	BookID int64   `json:"book_id"`
	Rating float32 `json:"rating"`
	Count  int64   `json:"count"`
}
