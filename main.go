package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	srv.ListenAndServe()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
