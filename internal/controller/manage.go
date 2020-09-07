package controller

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/bst27/plaudern/internal/database"
	"github.com/bst27/plaudern/internal/model/api/request"
	"github.com/bst27/plaudern/internal/model/api/response"
	"github.com/bst27/plaudern/internal/model/comment"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func registerManageRoutes(r *gin.Engine, config *configuration.Config, policy *bluemonday.Policy) {
	execFile, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	manage := r.Group("/manage")

	// TODO: Redirect URLs to allow Angular router to handle the  routing.
	// Otherwise calling URLs like this http://localhost:8080/manage/app/comments/2deac5b6-73ef-4fa6-b207-6fc5370e1e40
	// directly will not work because the server searches  for the file and cannot find it.
	manage.Static("/app", filepath.Join(filepath.Dir(execFile), "web"))

	manage.GET("/comment", func(ctx *gin.Context) {
		db := database.Get()

		cmnts, err := comment.GetAll(db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, response.NewGetComments(cmnts, policy))
	})

	manage.PUT("/comment/:commentId", func(ctx *gin.Context) {
		req, err := request.ParsePutComment(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		db := database.Get()
		cmnt, err := comment.GetOne(req.CommentId, db)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		if req.UpdateStatus {
			switch req.Status {
			case comment.STATUS_PUBLISHED:
				if err = cmnt.Approve(); err != nil {
					log.Println(err)
					ctx.JSON(http.StatusInternalServerError, gin.H{})
					return
				}
			case comment.STATUS_CREATED:
				if err = cmnt.Revoke(); err != nil {
					log.Println(err)
					ctx.JSON(http.StatusInternalServerError, gin.H{})
					return
				}
			default:
				ctx.JSON(http.StatusNotImplemented, gin.H{})
			}
		}

		if err = comment.Save(cmnt, db); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, response.NewPutComment(cmnt, policy))
	})
}
