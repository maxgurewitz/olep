package olep

import (
	"github.com/gin-gonic/gin"
)

// InitRequestHandler starts server responsible for handling payment requests.
func InitRequestHandler() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{

		v1.GET("/healthz", func(c *gin.Context) {
			c.String(200, "OK")
		})

		v1.POST("/payments", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

	}

	r.Run()
}
