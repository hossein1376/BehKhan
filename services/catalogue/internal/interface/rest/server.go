package rest

import (
	"net/http"
)

func Start(addr string) error {
	return http.ListenAndServe(addr, nil)
}
