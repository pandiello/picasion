package main

import (
	"log"
	"net/http"
	"time"
)

func LogWrapper(inner http.HandlerFunc, name string) http.HandlerFunc {
	return http.HandlerFunc(func(writter http.ResponseWriter, request *http.Request){
		
		start := time.Now()
		inner.ServeHTTP(writter, request)

		log.Printf(
			"%s\t%s\t%s\t%s",
			request.Method,
			request.RequestURI,
			name,
			time.Since(start),
		)
	})
}