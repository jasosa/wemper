package invitations

import (
	"fmt"
)

//UserNotFoundError represents a user not found error
type UserNotFoundError struct {
	UserID    string
	BaseError error
}

func (e *UserNotFoundError) Error() string {
	if e == nil {
		return ""
	}
	errorString := fmt.Sprintf("User not found. \"UserID\":{\"%s\"} \"Base Errors\":{\"%s\"}", e.UserID, e.BaseError.Error())
	return errorString
}

//ActionNotAllowedToUserError is returned when the user does not hav permission to perform an action
type ActionNotAllowedToUserError struct {
	UserID string
	Action string
}

func (e *ActionNotAllowedToUserError) Error() string {
	if e == nil {
		return ""
	}
	errorString := fmt.Sprintf("Action is not alllowed. \"UserID\":{\"%s\"} \"Action\":{\"%s\"}", e.UserID, e.Action)
	return errorString
}

//UserCouldNotBeAddedError is returned when there is an error try to adding a new user to the system
type UserCouldNotBeAddedError struct {
	BaseError error
}

func (e *UserCouldNotBeAddedError) Error() string {
	if e == nil {
		return ""
	}
	errorString := fmt.Sprintf("User was not added to the system. \"Base Errors\":{\"%s\"}", e.BaseError.Error())
	return errorString
}

//UserNotValidError is returned when the information about the user is not valid
type UserNotValidError struct {
	BaseError error
}

func (e *UserNotValidError) Error() string {
	if e == nil {
		return ""
	}
	errorString := fmt.Sprintf("User not valid. \"Base Errors\":{\"%s\"}", e.BaseError.Error())
	return errorString
}
