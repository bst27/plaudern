package controller

import (
	"fmt"
	"github.com/bst27/plaudern/internal/model/api/comment"
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

	r.POST("/comment", func(c *gin.Context) {
		newComment, err := comment.FromRequest(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		fmt.Println(newComment) // TODO: Persist comment

		c.JSON(http.StatusOK, gin.H{})
	})
}
