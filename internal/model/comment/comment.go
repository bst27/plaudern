package comment

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	id       string
	created  time.Time
	threadId string
	message  string
	author   string
}

func New(msg string, threadId string, author string) (*Comment, error) {
	if !IsValidMessage(msg) {
		return nil, errors.New("invalid message")
	}

	return &Comment{
		id:       uuid.New().String(),
		created:  time.Now(),
		threadId: threadId,
		message:  msg,
		author:   author,
	}, nil
}

func (r *Comment) GetFrontendData() map[string]interface{} {
	fd := make(map[string]interface{})
	fd["Id"] = r.id
	fd["Created"] = r.created
	fd["ThreadId"] = r.threadId
	fd["Message"] = r.message
	fd["Author"] = r.author
	return fd
}

func load(id string, created time.Time, threadId string, message string, author string) *Comment {
	return &Comment{
		id:       id,
		created:  created,
		threadId: threadId,
		message:  message,
		author:   author,
	}
}

func IsValidMessage(message string) bool {
	return message != ""
}
