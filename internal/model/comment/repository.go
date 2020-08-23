package comment

import (
	"database/sql"
	"time"
)

func Save(comment *Comment, db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO comments (id, created, message) VALUES (?, ?, ?);",
		comment.id,
		comment.created.Format(time.RFC3339),
		comment.message,
	)

	return err
}
