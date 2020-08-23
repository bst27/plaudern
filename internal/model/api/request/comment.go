package request

import (
	"errors"
	"github.com/bst27/plaudern/internal/model/comment"
	"github.com/gin-gonic/gin"
)

type CreateComment struct {
	Message  string
	ThreadId string
}

func ParseComment(r *gin.Context) (*CreateComment, error) {
	msg := r.PostForm("message")
	threadId := r.PostForm("threadId")

	if !comment.IsValidMessage(msg) {
		return nil, errors.New("invalid request message")
	}

	if threadId == "" {
		return nil, errors.New("invalid thread id")
	}

	return &CreateComment{
		Message:  msg,
		ThreadId: threadId,
	}, nil
}
