package comment

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	id      string
	created time.Time
	message string
}

func New(msg string) (*Comment, error) {
	if !IsValidMessage(msg) {
		return nil, errors.New("invalid message")
	}

	return &Comment{
		id:      uuid.New().String(),
		created: time.Now(),
		message: msg,
	}, nil
}

func load(id string, created time.Time, message string) *Comment {
	return &Comment{
		id:      id,
		created: created,
		message: message,
	}
}

func IsValidMessage(message string) bool {
	return message != ""
}
