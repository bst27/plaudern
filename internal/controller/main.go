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
	req := request.ParseGetAuth(ctx)

	if !ts.CheckToken(req.Token) {
		ctx.JSON(http.StatusForbidden, gin.H{})
	}

	return true
}

func checkXsrf(ctx *gin.Context) bool {
	auth := request.ParseGetAuth(ctx)
	xsrf := request.ParseXsrf(ctx)

	// Always block empty xsrf tokens
	if xsrf.Token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return false
	}

	// The xsrf token has to match with the first 32 characters of the auth token
	if len([]rune(auth.Token)) < 32 || xsrf.Token != string([]rune(auth.Token)[0:32]) {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return false
	}

	return true
}
