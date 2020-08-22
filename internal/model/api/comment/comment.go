package comment

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateComment struct {
	Created time.Time
	Message string
}

func FromRequest(r *gin.Context) (*CreateComment, error) {
	msg := r.PostForm("message")

	if msg == "" {
		return nil, errors.New("invalid comment message")
	}

	return &CreateComment{
		Created: time.Now(),
		Message: msg,
	}, nil
}
