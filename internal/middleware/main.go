package middleware

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Register(r *gin.Engine, config *configuration.Config) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:           config.AllowedOrigins,
		AllowOriginFunc:        nil,
		AllowMethods:           []string{"GET", "POST"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       false,
		ExposeHeaders:          nil,
		MaxAge:                 24 * time.Hour,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))
}
