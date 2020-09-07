package comment

import (
	"database/sql"
	"errors"
	"time"
)

func Save(comment *Comment, db *sql.DB) error {
	exists, err := exists(comment.id, db)
	if err != nil {
		return err
	}

	if exists {
		_, err = db.Exec(
			"UPDATE comments SET threadId = ?, message = ?, author = ?, status = ? WHERE id = ?;",
			comment.threadId,
			comment.message,
			comment.author,
			comment.status,
			comment.id,
		)
	} else {
		_, err = db.Exec(
			"INSERT INTO comments (id, created, threadId, message, author, status) VALUES (?, ?, ?, ?, ?, ?);",
			comment.id,
			comment.created.Format(time.RFC3339),
			comment.threadId,
			comment.message,
			comment.author,
			comment.status,
		)
	}

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

func exists(commentId string, db *sql.DB) (bool, error) {
	rows, err := db.Query(
		"SELECT id FROM comments WHERE id = ? ORDER BY created DESC LIMIT 1",
		commentId,
	)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	exists := rows.Next()
	if err := rows.Err(); err != nil {
		return false, err
	}

	return exists, nil
}

func GetOne(commentId string, db *sql.DB) (*Comment, error) {
	rows, err := db.Query(
		"SELECT id, created, threadId, message, author, status FROM comments WHERE id = ? ORDER BY created DESC LIMIT 1",
		commentId,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readRow(rows)
}

func readRow(rows *sql.Rows) (*Comment, error) {
	if !rows.Next() {
		return nil, errors.New("no row exists")
	}

	var id, created, threadId, message, author, status string
	if err := rows.Scan(&id, &created, &threadId, &message, &author, &status); err != nil {
		return nil, err
	}

	t, err := time.Parse(time.RFC3339, created)
	if err != nil {
		return nil, err
	}

	comment := load(id, t, threadId, message, author, status)

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comment, nil
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
