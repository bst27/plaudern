package controller

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/bst27/plaudern/internal/model/api/request"
	"github.com/bst27/plaudern/internal/model/auth"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

func RegisterRoutes(r *gin.Engine, config *configuration.Config) {
	policy := bluemonday.StrictPolicy()

	registerPublicRoutes(r, config, policy)
	registerManageRoutes(r, config, policy)
}

func checkAuth(ctx *gin.Context, ts *auth.TokenStore) bool {
	req, err := request.ParseGetAuth(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return false
	}

	if !ts.CheckToken(req.Token) {
		ctx.JSON(http.StatusForbidden, gin.H{})
	}

	return true
}
