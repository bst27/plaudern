package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbInstance *sql.DB
)

func Open() (*sql.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	db, err := sql.Open("sqlite3", ":memory:") // TODO: Use persistent database
	if err != nil {
		return nil, err
	}

	setup(db)

	dbInstance = db
	return Open()
}

func setup(db *sql.DB) error {
	sqlStmt := `
		CREATE TABLE comments (
			id TEXT PRIMARY KEY NOT NULL,
			created TEXT NOT NULL,
			threadId TEXT NOT NULL,
			message TEXT NOT NULL
		);
	`
	_, err := db.Exec(sqlStmt)
	return err
}
