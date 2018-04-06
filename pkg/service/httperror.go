package service

import (
	"github.com/jasosa/wemper/pkg/invitations"
	"github.com/jasosa/wemper/pkg/invitations/mysql"
	"net/http"
)

//HTTPError represents an error sent to the API customers
type HTTPError struct {
	HTTPStatusCode int
	ErrorMessage   string
}

func getHTTPErrorFromDomainError(err error) HTTPError {
	switch err.(type) {
	case *mysql.SQLOpeningDBError:
		return HTTPError{HTTPStatusCode: http.StatusServiceUnavailable, ErrorMessage: err.Error()}
	case *mysql.SQLQueryError:
		return HTTPError{HTTPStatusCode: http.StatusInternalServerError, ErrorMessage: err.Error()}

	case *invitations.UserNotFoundError:
		return HTTPError{HTTPStatusCode: http.StatusNotFound, ErrorMessage: err.Error()}
	case *invitations.UserNotValidError:
		return HTTPError{HTTPStatusCode: http.StatusBadRequest, ErrorMessage: err.Error()}
	case *invitations.UserCouldNotBeAddedError:
		return HTTPError{HTTPStatusCode: http.StatusInternalServerError, ErrorMessage: err.Error()}
	case *invitations.ActionNotAllowedToUserError:
		return HTTPError{HTTPStatusCode: http.StatusForbidden, ErrorMessage: err.Error()}

	case *RequestDecodeError:
		return HTTPError{HTTPStatusCode: http.StatusBadRequest, ErrorMessage: err.Error()}
	default:
		return HTTPError{HTTPStatusCode: http.StatusInternalServerError, ErrorMessage: err.Error()}
	}
}

//Write Write the http error using the response writer
func (e *HTTPError) Write(w http.ResponseWriter) {
	http.Error(w, e.ErrorMessage, e.HTTPStatusCode)
}
