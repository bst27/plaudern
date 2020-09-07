package controller

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func RegisterRoutes(r *gin.Engine, config *configuration.Config) {
	policy := bluemonday.StrictPolicy()

	registerPublicRoutes(r, config, policy)
	registerManageRoutes(r, config, policy)
}
