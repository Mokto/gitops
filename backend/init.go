package backend

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Init the frontend and api backend
func Init() {
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
		InitAPI(api)
	}

	go initGithub()

	r.Run(":8000")
}
