package service

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

//LoggingMiddleware middleware that adds basic login to a handler
func LoggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.WithFields(log.Fields{
			"requestedUrl": r.RequestURI,
		}).Info("Handler request")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
