package controller

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/bst27/plaudern/internal/database"
	"github.com/bst27/plaudern/internal/model/api/request"
	"github.com/bst27/plaudern/internal/model/api/response"
	"github.com/bst27/plaudern/internal/model/auth"
	"github.com/bst27/plaudern/internal/model/comment"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func registerManageRoutes(r *gin.Engine, config *configuration.Config, policy *bluemonday.Policy) {
	execFile, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	tokenStore := auth.NewTokenStore()

	manage := r.Group("/manage")

	manage.Static("/app", filepath.Join(filepath.Dir(execFile), "web"))

	r.NoRoute(func(ctx *gin.Context) {
		// Redirect URLs to allow Angular router to handle the routing. Otherwise calling URLs
		// like this http://localhost:8080/manage/app/comments/2deac5b6-73ef-4fa6-b207-6fc5370e1e40
		// directly will not work because the server searches for the file and cannot find it.
		// Read more: https://angular.io/guide/deployment#server-configuration

		if strings.HasPrefix(ctx.Request.URL.Path, "/manage/app/") {
			ctx.File(filepath.Join(filepath.Dir(execFile), "web", "index.html"))
		}
	})

	manage.GET("/comment", func(ctx *gin.Context) {
		if !checkAuth(ctx, tokenStore) {
			return
		}

		db := database.Get()

		cmnts, err := comment.GetAll(db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, response.NewGetComments(cmnts, policy, true))
	})

	manage.PUT("/comment/:commentId", func(ctx *gin.Context) {
		if !checkAuth(ctx, tokenStore) || !checkXsrf(ctx) {
			return
		}

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

		ctx.JSON(http.StatusOK, response.NewPutComment(cmnt, policy, true))
	})

	manage.POST("/login", func(ctx *gin.Context) {
		req, err := request.ParseLogin(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		// Block login without (= with empty/default) password
		if req.Password == "" || req.Password != config.AdminPassword {
			ctx.JSON(http.StatusOK, response.NewGetAuth(false))
			return
		}

		authToken := tokenStore.NewToken()
		xsrfToken := string([]rune(authToken)[0:32])

		ctx.SetCookie("auth-token", authToken, 60*60*24*2, "", "", false, true)  //TODO
		ctx.SetCookie("XSRF-TOKEN", xsrfToken, 60*60*24*2, "", "", false, false) //TODO

		ctx.JSON(http.StatusOK, response.NewGetAuth(tokenStore.CheckToken(authToken)))
	})

	manage.POST("/logout", func(ctx *gin.Context) {
		if !checkXsrf(ctx) {
			return
		}

		req := request.ParseGetAuth(ctx)
		tokenStore.RemoveToken(req.Token)
		ctx.SetCookie("auth-token", "", -1, "", "", false, true)  //TODO
		ctx.SetCookie("XSRF-TOKEN", "", -1, "", "", false, false) //TODO

		ctx.JSON(http.StatusOK, response.NewGetAuth(tokenStore.CheckToken(req.Token)))
	})

	manage.GET("/auth", func(ctx *gin.Context) {
		req := request.ParseGetAuth(ctx)
		ctx.JSON(http.StatusOK, response.NewGetAuth(tokenStore.CheckToken(req.Token)))
	})
}
