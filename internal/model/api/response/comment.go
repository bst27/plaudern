package response

import (
	"github.com/bst27/plaudern/internal/model/comment"
	"github.com/microcosm-cc/bluemonday"
)

type GetComments struct {
	Comments []map[string]interface{}
}

type PutComment struct {
	Comment interface{}
}

type GetAuth struct {
	Authorized bool
}

func NewGetComments(
	comments []*comment.Comment,
	policy *bluemonday.Policy,
	asAdmin bool,
) *GetComments {
	gc := &GetComments{}
	gc.Comments = make([]map[string]interface{}, 0)

	for _, cmnt := range comments {
		gc.Comments = append(gc.Comments, cmnt.GetFrontendData(policy, asAdmin))
	}

	return gc
}

func NewPutComment(
	cmnt *comment.Comment,
	policy *bluemonday.Policy,
	asAdmin bool,
) *PutComment {
	return &PutComment{
		Comment: cmnt.GetFrontendData(policy, asAdmin),
	}
}

func NewGetAuth(authorized bool) *GetAuth {
	return &GetAuth{
		Authorized: authorized,
	}
}
