package controller

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/bst27/plaudern/internal/database"
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
}
