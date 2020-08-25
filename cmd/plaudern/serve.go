package main

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/bst27/plaudern/internal/controller"
	"github.com/bst27/plaudern/internal/database"
	"github.com/bst27/plaudern/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start webserver",
		Long:  "Start webserver to serve and receive comments.",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := configuration.ReadFile(cmd.Flag("config").Value.String())
			if err != nil {
				log.Fatalln(err)
			}

			// Open database at the beginning for initialization; subsequent calls can use database.Get() instead.
			_, err = database.Open(config)
			if err != nil {
				log.Fatalln(err)
			}

			r := gin.Default()

			middleware.Register(r, config)
			controller.RegisterRoutes(r)

			r.Run(":" + strconv.FormatInt(config.Port, 10))
		},
	}
)

func init() {
	serveCmd.Flags().String("config", "plaudern-config.json", "Config file to use")
	err := serveCmd.MarkFlagRequired("config")
	if err != nil {
		log.Fatalln(err)
	}

	rootCmd.AddCommand(serveCmd)
}
