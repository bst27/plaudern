package database

import (
	"database/sql"
	"github.com/bst27/plaudern/internal/configuration"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	dbInstance *sql.DB
)

// Use this function to open the database exactly once
func Open(config *configuration.Config) (*sql.DB, error) {
	if dbInstance != nil {
		log.Fatalln("Database already open")
	}

	db, err := sql.Open("sqlite3", config.DatabaseFile)
	if err != nil {
		return nil, err
	}

	setup(db)

	dbInstance = db
	return Get(), nil
}

// Only use this function after database has successfully been opened with Open()
func Get() *sql.DB {
	return dbInstance
}

func setup(db *sql.DB) error {
	sqlStmt := `
		CREATE TABLE comments (
			id TEXT PRIMARY KEY NOT NULL,
			created TEXT NOT NULL,
			threadId TEXT NOT NULL,
			message TEXT NOT NULL,
			author TEXT NOT NULL
		);
	`
	_, err := db.Exec(sqlStmt)
	return err
}
