package mysql

import (
	"fmt"
)

//SQLQueryError represents an SQL error
type SQLQueryError struct {
	Query     string
	BaseError error
}

func (e *SQLQueryError) Error() string {
	if e == nil {
		return ""
	}
	errorString := fmt.Sprintf("SQL query failed. \"Query\":{\"%s\"}, \"Errors\":{\"%s\"}", e.Query, e.BaseError.Error())
	return errorString
}

//SQLOpeningDBError represents an SQL error
type SQLOpeningDBError struct {
	BaseError error
}

func (e *SQLOpeningDBError) Error() string {
	if e == nil {
		return ""
	}
	errorString := fmt.Sprintf("Open db failed. \"Errors\":{\"%s\"}", e.BaseError.Error())
	return errorString
}
