package controller

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/bst27/plaudern/internal/database"
	"github.com/bst27/plaudern/internal/model/api/request"
	"github.com/bst27/plaudern/internal/model/api/response"
	"github.com/bst27/plaudern/internal/model/comment"
	"github.com/bst27/plaudern/internal/webhook"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
	"time"
)

func RegisterRoutes(r *gin.Engine, config *configuration.Config) {
	policy := bluemonday.StrictPolicy()

	registerManageRoutes(r, config, policy)

	r.GET("/ping", func(c *gin.Context) {
		db := database.Get()

		if err := db.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Time":  time.Now().Unix(),
				"Error": "Database not available",
			})
			return
		}

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

		cmnt, err := comment.New(req.Message, req.ThreadId, req.Author)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		db := database.Get()

		err = comment.Save(cmnt, db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		if config.NewCommentWebhook != "" {
			go func() {
				err := webhook.New().Receive(config.NewCommentWebhook, gin.H{
					"Author":   req.Author,
					"ThreadId": req.ThreadId,
					"Message":  req.Message,
				})

				if err != nil {
					log.Println(err)
				}
			}()
		}

		ctx.JSON(http.StatusOK, gin.H{})
	})

	r.GET("/comment", func(ctx *gin.Context) {
		req, err := request.ParseGetComments(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		db := database.Get()

		cmnts, err := comment.GetByThread(req.ThreadId, db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, response.NewGetComments(cmnts, policy))
	})

	r.GET("/thread/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{}) // TODO: Implement
	})
}
