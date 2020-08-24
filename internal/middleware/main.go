package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Register(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:           []string{"http://localhost:8083"}, // TODO: Make configurable
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
