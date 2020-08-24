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

type GetComments struct {
	ThreadId string
}

func ParseCreateComment(r *gin.Context) (*CreateComment, error) {
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

func ParseGetComments(r *gin.Context) (*GetComments, error) {
	threadId := r.Query("threadId")

	if threadId == "" {
		return nil, errors.New("invalid thread id")
	}

	return &GetComments{
		ThreadId: threadId,
	}, nil
}
