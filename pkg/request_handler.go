package olep

import (
	"github.com/gin-gonic/gin"
)

// InitRequestHandler starts server responsible for handling payment requests.
func InitRequestHandler() {
	r := gin.Default()

	r.POST("/api/v1/payments", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
