package comment

import (
	"errors"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"time"
)

const (
	STATUS_CREATED   = "created"
	STATUS_PUBLISHED = "published"
	STATUS_DELETED   = "deleted"
)

type Comment struct {
	id       string
	created  time.Time
	threadId string
	message  string
	author   string
	status   string
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
		status:   STATUS_CREATED,
	}, nil
}

func (r *Comment) GetFrontendData(policy *bluemonday.Policy) map[string]interface{} {
	// When changing the response make sure to also update the corresponding typescript interface:
	// web/src/app/models/comment.ts

	fd := make(map[string]interface{})
	fd["Id"] = r.id
	fd["Created"] = r.created
	fd["ThreadId"] = policy.Sanitize(r.threadId)
	fd["ThreadIdInsecure"] = r.threadId
	fd["Message"] = policy.Sanitize(r.message)
	fd["MessageInsecure"] = r.message
	fd["Author"] = policy.Sanitize(r.author)
	fd["AuthorInsecure"] = r.author
	fd["Status"] = r.status
	return fd
}

func (r *Comment) Approve() error {
	switch r.status {
	case STATUS_CREATED:
		r.status = STATUS_PUBLISHED
		return nil
	case STATUS_PUBLISHED:
		// nothing to do
		return nil
	default:
		return errors.New("invalid status change")
	}
}

func (r *Comment) Revoke() error {
	switch r.status {
	case STATUS_CREATED:
		// nothing to do
		return nil
	case STATUS_PUBLISHED:
		r.status = STATUS_CREATED
		return nil
	default:
		return errors.New("invalid status change")
	}
}

func load(id string, created time.Time, threadId string, message string, author string, status string) *Comment {
	return &Comment{
		id:       id,
		created:  created,
		threadId: threadId,
		message:  message,
		author:   author,
		status:   status,
	}
}

func IsValidMessage(message string) bool {
	return message != ""
}
