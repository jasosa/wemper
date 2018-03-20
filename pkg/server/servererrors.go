package server

import (
	"fmt"
)

//RequestDecodeError represents a decoding json error
type RequestDecodeError struct {
	BaseError error
}

func (e *RequestDecodeError) Error() string {
	if e == nil {
		return ""
	}
	errorString := fmt.Sprintf("Error decoding JSON request. \"Errors\":{\"%s\"}", e.BaseError.Error())
	return errorString
}
