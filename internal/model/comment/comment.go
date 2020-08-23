package comment

import (
	"errors"
	"time"
)

type Comment struct {
	created time.Time
	message string
}

func New(msg string) (*Comment, error) {
	if !IsValidMessage(msg) {
		return nil, errors.New("invalid message")
	}

	return &Comment{
		created: time.Now(),
		message: msg,
	}, nil
}

func IsValidMessage(message string) bool {
	return message != ""
}
