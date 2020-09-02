package response

import (
	"github.com/bst27/plaudern/internal/model/comment"
	"github.com/microcosm-cc/bluemonday"
)

type GetComments struct {
	Comments []map[string]interface{}
}

func NewGetComments(comments []*comment.Comment, policy *bluemonday.Policy) *GetComments {
	gc := &GetComments{}
	gc.Comments = make([]map[string]interface{}, 0)

	for _, cmnt := range comments {
		gc.Comments = append(gc.Comments, cmnt.GetFrontendData(policy))
	}

	return gc
}
