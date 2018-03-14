package server

import "net/http"

//AppHandler ...
type AppHandler func(http.ResponseWriter, *http.Request) error

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		httperror := getHTTPError(err)
		http.Error(w, httperror.ErrorMessage, httperror.HTTPStatusCode)
	}
}
