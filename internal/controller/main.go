package controller

import (
	"github.com/bst27/plaudern/internal/database"
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
		req, err := request.ParseCreateComment(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		cmnt, err := comment.New(req.Message, req.ThreadId)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		db, err := database.Open()
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		err = comment.Save(cmnt, db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	})

	r.GET("/thread/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{}) // TODO: Implement
	})
}
