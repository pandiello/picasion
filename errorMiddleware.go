package main

import(
	"net/http"
)

func(eh httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := eh(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}