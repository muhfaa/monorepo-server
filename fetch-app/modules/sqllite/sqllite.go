package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewDatabaseConnection(dbName string) (*sql.DB, error) {
	database := fmt.Sprintf("../%v", dbName)
	sqlliteDatabase, err := sql.Open("sqlite3", database)
	if err != nil {
		return nil, err
	}
	return sqlliteDatabase, nil
}
