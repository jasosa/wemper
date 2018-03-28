package service

import (
	"log"
	"net/http"
)

//HTTPErrorMiddleware gets any error from the decorated handler and converts it into an ClientAPIError
type HTTPErrorMiddleware func(http.ResponseWriter, *http.Request) error

func (fn HTTPErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		httperror := getClientAPIError(err)
		http.Error(w, httperror.ErrorMessage, httperror.HTTPStatusCode)
	}
}

//LoggingMiddleware middleware that adds basic login to a handler
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
