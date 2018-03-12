package mysql

import (
	"database/sql"
	//ok
	_ "github.com/go-sql-driver/mysql"
)

//Connection ...
type Connection interface {
	OpenConnection(stringConn string) (*sql.DB, error)
}

//DbConnection encapsulates the opening of an SQL connection
type DbConnection struct {
}

//OpenConnection open an sql connection
func (db *DbConnection) OpenConnection(stringConn string) (*sql.DB, error) {
	return sql.Open("mysql", stringConn)
}
