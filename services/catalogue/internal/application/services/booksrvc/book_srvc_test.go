package booksrvc

import (
	"context"
	"reflect"
	"testing"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/repository"
	"github.com/hossein1376/BehKhan/catalogue/test/mocks"
)

func TestBookSrvc_GetByID(t *testing.T) {
	tests := []struct {
		name    string
		db      repository.Pool
		id      entity.BookID
		want    *entity.Book
		wantErr bool
	}{
		{
			name:    "Get Book",
			db:      mocks.Pool{},
			id:      entity.BookID(1),
			want:    &entity.Book{ID: entity.BookID(1), Title: "1"},
			wantErr: false,
		},
		{
			name:    "Book Not Found",
			db:      mocks.Pool{},
			id:      entity.BookID(0),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := BookSrvc{
				db: tt.db,
			}
			got, err := c.GetByID(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
