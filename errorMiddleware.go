package main

import(
	"net/http"
)

type errorMiddlewareAdapter func(w http.ResponseWriter, r *http.Request) error

func errorMiddleware(eh httpHandler) http.Handler {
	return errorMiddlewareAdapter(eh)
}

func(eh errorMiddlewareAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := eh(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}