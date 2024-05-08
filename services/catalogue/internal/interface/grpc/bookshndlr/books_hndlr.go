package bookshndlr

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/service"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/grpc/pb/cataloguePB"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/grpc/serde"
)

type BooksHndlr struct {
	Services service.Service
	Logger   *slog.Logger

	cataloguePB.UnimplementedBookServiceServer
}

func New(srv *grpc.Server, srvc service.Service, logger *slog.Logger) BooksHndlr {
	bookHandlers := BooksHndlr{
		Services: srvc,
		Logger:   logger,
	}

	cataloguePB.RegisterBookServiceServer(srv, bookHandlers)

	return bookHandlers
}

func (h BooksHndlr) GetBook(ctx context.Context, request *cataloguePB.BookRequest) (*cataloguePB.BookResponse, error) {
	id := request.GetId()
	if id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "id must be positive")
	}

	book, err := h.Services.BookSrvc.GetByID(ctx, entity.BookID(id))
	if err != nil {
		return nil, serde.Code(err)
	}

	return &cataloguePB.BookResponse{
		Books: &cataloguePB.Book{
			Id:   book.ID.ToInt64(),
			Name: book.Title,
		},
	}, nil
}
