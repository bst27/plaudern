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

func GetAll(db *sql.DB) ([]*Comment, error) {
	rows, err := db.Query("SELECT id, created, message FROM comments ORDER BY created DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		var id, created, message string
		if err = rows.Scan(&id, &created, &message); err != nil {
			return nil, err
		}

		t, err := time.Parse(time.RFC3339, created)
		if err != nil {
			return nil, err
		}

		comment := load(id, t, message)

		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
