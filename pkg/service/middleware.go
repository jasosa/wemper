package service

import "net/http"

//APIErrorMiddleware gets any error from the decorated handler and converts it into an ClientAPIError
type APIErrorMiddleware func(http.ResponseWriter, *http.Request) error

func (fn APIErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		httperror := getClientAPIError(err)
		http.Error(w, httperror.ErrorMessage, httperror.HTTPStatusCode)
	}
}
