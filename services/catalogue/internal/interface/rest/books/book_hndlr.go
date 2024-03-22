package books

import (
	"net/http"
)

type BookRestHndlr struct {
}

func NewBookRestHndlr() BookRestHndlr {
	return BookRestHndlr{}
}

func (h *BookRestHndlr) CreateNewBookHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *BookRestHndlr) GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {
}
