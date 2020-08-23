package request

import (
	"errors"
	"github.com/bst27/plaudern/internal/model/comment"
	"github.com/gin-gonic/gin"
)

type CreateComment struct {
	Message string
}

func ParseComment(r *gin.Context) (*CreateComment, error) {
	msg := r.PostForm("message")

	if !comment.IsValidMessage(msg) {
		return nil, errors.New("invalid request message")
	}

	return &CreateComment{
		Message: msg,
	}, nil
}
