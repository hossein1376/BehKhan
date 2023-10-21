package handlers

import (
	"context"

	"github.com/hossein1376/BehKhan/catalogue/proto/cataloguePB"
)

type Server struct {
	cataloguePB.UnimplementedBookServiceServer
}

func (s Server) GetBook(_ context.Context, in *cataloguePB.BookRequest) (*cataloguePB.BookResponse, error) {
	books := make([]*cataloguePB.Book, 0, len(in.GetId()))

	return &cataloguePB.BookResponse{
		Books: books,
	}, nil
}
