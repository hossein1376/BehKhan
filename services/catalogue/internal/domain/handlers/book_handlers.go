package handlers

import (
	"net/http"
)

type BookHandlers interface {
	CreateNewBookHandler(w http.ResponseWriter, r *http.Request)
	GetBookByIDHandler(w http.ResponseWriter, r *http.Request)
}
