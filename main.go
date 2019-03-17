package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"gitops/backend"
	"os"
)

func main() {
	fmt.Println("Run server")
	r := gin.Default()

	if os.Getenv("GIN_MODE") == "release" {
		r.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))
		r.NoRoute(func(c *gin.Context) {
			c.File("frontend/build/index.html")
		})
	}

	api := r.Group("/api")
	{
		backend.InitAPI(api)
	}

	go backend.Init()

	r.Run(":8000")
}
