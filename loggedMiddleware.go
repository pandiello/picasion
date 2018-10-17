package main

import (
	"log"
	"net/http"
	"time"
)

func LogMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(writter http.ResponseWriter, request *http.Request) {

		start := time.Now()
		inner.ServeHTTP(writter, request)

		log.Printf(
			"%s\t%s\t%s",
			request.Method,
			request.RequestURI,
			time.Since(start),
		)
	})
}
