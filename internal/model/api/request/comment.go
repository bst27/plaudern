package request

import (
	"errors"
	"github.com/bst27/plaudern/internal/model/comment"
	"github.com/gin-gonic/gin"
)

type CreateComment struct {
	Author   string
	Message  string
	ThreadId string
}

type GetComments struct {
	ThreadId string
}

type PutComment struct {
	CommentId    string
	UpdateStatus bool
	Status       string
}

type Login struct {
	Password string
}

type Auth struct {
	Token string
}

type Xsrf struct {
	Token string
}

func ParseCreateComment(r *gin.Context) (*CreateComment, error) {
	author := r.PostForm("author")
	msg := r.PostForm("message")
	threadId := r.PostForm("threadId")

	if !comment.IsValidMessage(msg) {
		return nil, errors.New("invalid request message")
	}

	if threadId == "" {
		return nil, errors.New("invalid thread id")
	}

	return &CreateComment{
		Author:   author,
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

func ParsePutComment(r *gin.Context) (*PutComment, error) {
	commentId := r.Param("commentId")
	if commentId == "" {
		return nil, errors.New("invalid comment id")
	}

	updateStatus := false
	status := r.DefaultPostForm("status", "-1")
	if status != "-1" {
		switch status {
		case comment.STATUS_CREATED:
			updateStatus = true
		case comment.STATUS_PUBLISHED:
			updateStatus = true
		case comment.STATUS_DELETED:
			updateStatus = true
		default:
			return nil, errors.New("invalid comment status")
		}
	} else {
		status = ""
	}

	return &PutComment{
		CommentId:    commentId,
		UpdateStatus: updateStatus,
		Status:       status,
	}, nil
}

func ParseLogin(r *gin.Context) (*Login, error) {
	password := r.DefaultPostForm("password", "")

	return &Login{
		Password: password,
	}, nil
}

func ParseGetAuth(r *gin.Context) *Auth {
	token, err := r.Cookie("auth-token")

	if err != nil {
		// Return empty token if cookie is not present
		return &Auth{
			Token: "",
		}
	}

	return &Auth{
		Token: token,
	}
}

func ParseXsrf(r *gin.Context) *Xsrf {
	token, err := r.Cookie("XSRF-TOKEN")

	if err != nil {
		// Return empty token if cookie is not present
		return &Xsrf{
			Token: "",
		}
	}

	return &Xsrf{
		Token: token,
	}
}
