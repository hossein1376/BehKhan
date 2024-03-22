package handlers

import (
	"net/http"
)

type BookHandler interface {
	CreateNewBookHandler(w http.ResponseWriter, r *http.Request)
	GetBookByIDHandler(w http.ResponseWriter, r *http.Request)
}
