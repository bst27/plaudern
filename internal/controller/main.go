package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Time": time.Now().Unix(),
		})
	})
}
