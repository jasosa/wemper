package service

import (
	"github.com/jasosa/wemper/pkg/invitations"
	"github.com/jasosa/wemper/pkg/invitations/mysql"
)

//ClientAPIError represents an error sent to the client
type ClientAPIError struct {
	HTTPStatusCode int
	ErrorMessage   string
}

func getClientAPIError(err error) ClientAPIError {
	switch err.(type) {
	case *mysql.SQLOpeningDBError:
		return ClientAPIError{HTTPStatusCode: http503ServiceUnavailable, ErrorMessage: err.Error()}
	case *mysql.SQLQueryError:
		return ClientAPIError{HTTPStatusCode: http500InternalServerError, ErrorMessage: err.Error()}

	case *invitations.UserNotFoundError:
		return ClientAPIError{HTTPStatusCode: http404NotFound, ErrorMessage: err.Error()}
	case *invitations.UserNotValidError:
		return ClientAPIError{HTTPStatusCode: http400BadRequest, ErrorMessage: err.Error()}
	case *invitations.UserCouldNotBeAddedError:
		return ClientAPIError{HTTPStatusCode: http500InternalServerError, ErrorMessage: err.Error()}
	case *invitations.ActionNotAllowedToUserError:
		return ClientAPIError{HTTPStatusCode: http403Forbidden, ErrorMessage: err.Error()}

	case *RequestDecodeError:
		return ClientAPIError{HTTPStatusCode: http400BadRequest, ErrorMessage: err.Error()}
	default:
		return ClientAPIError{HTTPStatusCode: http500InternalServerError, ErrorMessage: err.Error()}
	}
}

var http400BadRequest = 400
var http403Forbidden = 403
var http404NotFound = 404
var http500InternalServerError = 500
var http503ServiceUnavailable = 503
