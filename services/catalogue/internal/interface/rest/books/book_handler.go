package books

import (
	"net/http"
)

type BookHandler struct {
}

func NewBookHandler() BookHandler {
	return BookHandler{}
}

func (h *BookHandler) CreateNewBookHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {

}
