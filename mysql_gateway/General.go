package mysql_gateway

import (
	"database/sql"
)

type Gateway struct {
}

var (
	ConnectionString = "rest:password@tcp(localhost:3306)/books"
)

func getDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}
