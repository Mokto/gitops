package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitAPI - the backend api
func InitAPI(api *gin.RouterGroup) {
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
