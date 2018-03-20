package service

import (
	"github.com/jasosa/wemper/pkg/invitations"
	"github.com/jasosa/wemper/pkg/invitations/mysql"
)

//HTTPError represents an error sent to the client
type HTTPError struct {
	HTTPStatusCode int
	ErrorMessage   string
}

func getHTTPError(err error) HTTPError {
	switch err.(type) {
	case *mysql.SQLOpeningDBError:
		return HTTPError{HTTPStatusCode: http503ServiceUnavailable, ErrorMessage: err.Error()}
	case *mysql.SQLQueryError:
		return HTTPError{HTTPStatusCode: http500InternalServerError, ErrorMessage: err.Error()}

	case *invitations.UserNotFoundError:
		return HTTPError{HTTPStatusCode: http404NotFound, ErrorMessage: err.Error()}
	case *invitations.UserNotValidError:
		return HTTPError{HTTPStatusCode: http400BadRequest, ErrorMessage: err.Error()}
	case *invitations.UserCouldNotBeAddedError:
		return HTTPError{HTTPStatusCode: http500InternalServerError, ErrorMessage: err.Error()}
	case *invitations.ActionNotAllowedToUserError:
		return HTTPError{HTTPStatusCode: http403Forbidden, ErrorMessage: err.Error()}

	case *RequestDecodeError:
		return HTTPError{HTTPStatusCode: http400BadRequest, ErrorMessage: err.Error()}
	default:
		return HTTPError{HTTPStatusCode: http500InternalServerError, ErrorMessage: err.Error()}
	}
}

var http400BadRequest = 400
var http403Forbidden = 403
var http404NotFound = 404
var http500InternalServerError = 500
var http503ServiceUnavailable = 503
