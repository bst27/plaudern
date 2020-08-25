package main

import (
	"github.com/bst27/plaudern/internal/configuration"
	"github.com/bst27/plaudern/internal/controller"
	"github.com/bst27/plaudern/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	config := configuration.GetDefault()
	r := gin.Default()

	middleware.Register(r)
	controller.RegisterRoutes(r)

	log.Println(r.Run(":" + strconv.FormatInt(config.Port, 10)))
}
