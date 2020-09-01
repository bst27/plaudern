package controller

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

func registerManageRoutes(r *gin.Engine, config *configuration.Config) {
	execFile, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	r.Static("/manage", filepath.Join(filepath.Dir(execFile), "web"))
}
