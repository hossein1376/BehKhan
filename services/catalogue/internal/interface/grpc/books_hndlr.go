package grpc

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/dto"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/services"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/grpc/pb/cataloguePB"
)

type BooksHndlr struct {
	Services services.Service
	Logger   *slog.Logger

	cataloguePB.UnimplementedBookServiceServer
}

func NewBooksHndlr(srv *grpc.Server, srvc services.Service, logger *slog.Logger) BooksHndlr {
	bookHandlers := BooksHndlr{
		Services: srvc,
		Logger:   logger,
	}

	cataloguePB.RegisterBookServiceServer(srv, bookHandlers)

	return bookHandlers
}

func (h BooksHndlr) GetBook(ctx context.Context, request *cataloguePB.BookRequest) (*cataloguePB.BookResponse, error) {
	req := &dto.GetBookByIDRequest{
		ID: request.GetId(),
	}
	book, err := h.Services.BookSrvc.GetByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return &cataloguePB.BookResponse{
		Books: &cataloguePB.Book{
			Id:   book.ID,
			Name: book.Name,
		},
	}, nil
}
