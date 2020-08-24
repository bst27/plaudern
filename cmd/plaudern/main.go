package main

import (
	"github.com/bst27/plaudern/internal/controller"
	"github.com/bst27/plaudern/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	middleware.Register(r)
	controller.RegisterRoutes(r)

	log.Println(r.Run(":8080"))
}
