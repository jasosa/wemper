package mysql

import (
	"database/sql"
	//ok
	_ "github.com/go-sql-driver/mysql"
)

//DbConnection is a wrapper of an SQL connection
type DbConnection struct {
}

//OpenConnection open an sql connection
func (db *DbConnection) OpenConnection(stringConn string) (*sql.DB, error) {
	return sql.Open("mysql", stringConn)
}

// MockDbConnection is a wrapper for a mock connection
type MockDbConnection struct {
	mockdb *sql.DB
}

//OpenConnection open an mock slq connection
func (db *MockDbConnection) OpenConnection(stringConn string) (*sql.DB, error) {
	return db.mockdb, db.mockdb.Ping()
}
