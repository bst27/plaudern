package main

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/bst27/plaudern/internal/controller"
	"github.com/bst27/plaudern/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start webserver",
		Long:  "Start webserver to serve and receive comments.",
		Run: func(cmd *cobra.Command, args []string) {
			config := configuration.GetDefault()
			r := gin.Default()

			middleware.Register(r)
			controller.RegisterRoutes(r)

			r.Run(":" + strconv.FormatInt(config.Port, 10))
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}
