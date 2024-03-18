package dto

type CreateBookRequest struct {
	Name   string
	Author string
}

type GetBookByIDRequest struct {
	ID int
}
