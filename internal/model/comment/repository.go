package comment

import (
	"database/sql"
	"time"
)

func Save(comment *Comment, db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO comments (id, created, threadId, message, author, status) VALUES (?, ?, ?, ?, ?, ?);",
		comment.id,
		comment.created.Format(time.RFC3339),
		comment.threadId,
		comment.message,
		comment.author,
		comment.status,
	)

	return err
}

func GetByThread(threadId string, db *sql.DB) ([]*Comment, error) {
	rows, err := db.Query("SELECT id, created, threadId, message, author, status FROM comments WHERE threadId = ? ORDER BY created DESC", threadId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readRows(rows)
}

func GetAll(db *sql.DB) ([]*Comment, error) {
	rows, err := db.Query("SELECT id, created, threadId, message, author, status FROM comments ORDER BY created DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readRows(rows)
}

func readRows(rows *sql.Rows) ([]*Comment, error) {
	var comments []*Comment
	for rows.Next() {
		var id, created, threadId, message, author, status string
		if err := rows.Scan(&id, &created, &threadId, &message, &author, &status); err != nil {
			return nil, err
		}

		t, err := time.Parse(time.RFC3339, created)
		if err != nil {
			return nil, err
		}

		comment := load(id, t, threadId, message, author, status)

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
