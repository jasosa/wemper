package service

import (
	"log"
	"net/http"
)

//APIErrorMiddleware gets any error from the decorated handler and converts it into an ClientAPIError
type APIErrorMiddleware func(http.ResponseWriter, *http.Request) error

func (fn APIErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		httperror := getClientAPIError(err)
		http.Error(w, httperror.ErrorMessage, httperror.HTTPStatusCode)
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
