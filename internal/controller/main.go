package controller

import (
	"fmt"
	"github.com/bst27/plaudern/internal/model/api/request"
	"github.com/bst27/plaudern/internal/model/comment"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Time": time.Now().Unix(),
		})
	})

	r.POST("/comment", func(ctx *gin.Context) {
		req, err := request.ParseComment(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		cmnt, err := comment.New(req.Message)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		fmt.Println(cmnt) // TODO: Persist request

		ctx.JSON(http.StatusOK, gin.H{})
	})

	r.GET("/thread/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{}) // TODO: Implement
	})
}
